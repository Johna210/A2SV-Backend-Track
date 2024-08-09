package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_Name *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name  *string            `json:"last_name" validate:"required,min=2,max=100"`
	User_Name  *string            `json:"user_name" validate:"required,min=5"`
	Email      *string            `json:"email" validate:"email,required"`
	Password   *string            `json:"password" validate:"required,min=6"`
	User_Role  *string            `json:"user_role" validate:"required,eq=ADMIN|eq=USER"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
