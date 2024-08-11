package usecases

import (
	"context"
	"time"

	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
	infrastructure "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Infrastructure"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) GetUserByUsername(c context.Context, userName string) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByUsername(ctx, userName)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return infrastructure.CreateAccessToken(user, secret, expiry)
}
