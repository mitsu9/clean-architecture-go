package usecase

import (
	"github.com/mitsu9/clean-architecture-go/internal/api/domain/repository"
)

type GetUserUsecase struct {
	UserRepository *repository.UserRepository
}

func NewGetUserUsecase(userRepository *repository.UserRepository) *GetUserUsecase {
	return &GetUserUsecase{
		UserRepository: userRepository,
	}
}

func (u *GetUserUsecase) Execute() error {
	// TODO: impl
	return nil
}
