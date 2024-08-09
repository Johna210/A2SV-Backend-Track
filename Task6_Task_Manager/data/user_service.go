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

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash Password")
	}

	return string(hashedPassword), nil
}

func comparePasswords(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

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
