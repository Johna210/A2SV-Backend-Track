package usecases

import (
	"context"
	"time"

	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
)

type promoteUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewPromoteUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &promoteUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *promoteUsecase) Promote(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.Promote(ctx, id)
}
