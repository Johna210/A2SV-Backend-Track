package data

import (
	"context"
	"errors"
	"time"

	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User = models.User

var jwtSecret = []byte("JWT_SECRET_KEY")

const tokenExpirationDuration = time.Hour * 24

// hashPassword takes a password string and returns the hashed version of the password.
// It uses bcrypt.GenerateFromPassword to generate the hash with the default cost.
// If an error occurs during the hashing process, it returns an empty string and an error.
// Otherwise, it returns the hashed password as a string and a nil error.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash Password")
	}

	return string(hashedPassword), nil
}

// comparePasswords compares a plain text password with a hashed password and returns an error if they do not match.
// It uses bcrypt.CompareHashAndPassword to perform the comparison.
//
// Parameters:
// - password: The plain text password to compare.
// - hashedPassword: The hashed password to compare against.
//
// Returns:
// - error: An error if the passwords do not match, nil otherwise.
func comparePasswords(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

// Register registers a new user in the system.
// It checks if the email and username are already taken.
// If there are no existing users, the registered user is assigned the role of an admin.
// The user's password is hashed before storing it in the database.
// The created user object is then inserted into the UserCollection.
// If any error occurs during the registration process, an error is returned.
// Otherwise, the registered user is returned along with a nil error.
func Register(user User) (User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Check if email already taken
	count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return User{}, errors.New("error occured while checking for email")
	}

	if count > 0 {
		return User{}, errors.New("email already in use")
	}

	// Check if username already taken
	count, err = UserCollection.CountDocuments(ctx, bson.M{"user_name": user.User_Name})
	if err != nil {
		return User{}, errors.New("error occured while checking for email")
	}

	if count > 0 {
		return User{}, errors.New("userName already taken")
	}

	// Check if there is no any user in the data base and make the user admin else normal user.
	count, err = UserCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return User{}, errors.New("error occured while finding users")
	}

	// instantiate user_role to be User
	userRole := "USER"
	user.User_Role = &userRole

	// make user_role to be ADMIN if there is no user exists
	if count == 0 {
		adminRole := "ADMIN"
		user.User_Role = &adminRole
	}

	hashedPassword, err := hashPassword(*user.Password)

	if err != nil {
		return User{}, errors.New("unable to hash password")
	}
	// Create a user object for the db
	user.Password = &hashedPassword
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()

	_, insertErr := UserCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return User{}, errors.New("user was not created")
	}

	return user, nil

}

// Login authenticates a user by their username and password and generates a JWT token upon successful authentication.
// It takes the user_name and password as input parameters and returns the generated JWT token as a string and an error, if any.
// If the user is not found, it returns an error with the message "user not found".
// If the username or password is incorrect, it returns an error with the message "incorrect username or password".
// If there is an error while generating the token, it returns an error with the message "unable to generate token".
func Login(user_name, password string) (string, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user User
	err := UserCollection.FindOne(ctx, bson.M{"user_name": user_name}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = comparePasswords(password, *user.Password)

	if err != nil {
		return "", errors.New("incorrect username or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.User_Name,
		"role":      user.User_Role,
		"exp":       time.Now().Add(tokenExpirationDuration).Unix(),
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("unable to generate token")
	}

	return jwtToken, nil
}

// Promote promotes a user to an admin role.
// It takes the user's ID as a parameter and updates the user's role to "ADMIN" in the database.
// If the user ID is invalid or there is an error in updating the user's role, an error is returned.
func Promote(id string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ObjectId, err := ChangeIdToObjectId(id)
	if err != nil {
		return err
	}

	var user User
	err = UserCollection.FindOne(ctx, bson.M{"_id": ObjectId}).Decode(&user)

	if err != nil {
		return err
	}

	// Update the user's role to ADMIN
	var adminRole = "ADMIN"
	user.User_Role = &adminRole

	// Save the updated user back to the database
	_, err = UserCollection.UpdateOne(
		ctx,
		bson.M{"_id": ObjectId},
		bson.M{"$set": bson.M{"user_role": user.User_Role}},
	)
	if err != nil {
		return err
	}

	return nil

}
