package domain

import "context"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (user User, err error)
	GetUserByUsername(c context.Context, userName string) (user User, err error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
}
