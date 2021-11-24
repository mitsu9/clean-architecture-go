package dynamodb

import (
	"github.com/guregu/dynamo"
	"github.com/mitsu9/clean-architecture-go/internal/api/domain/repository"
)

type UserRepository struct {
	DB *dynamo.DB
}

func NewUserRepository(db *dynamo.DB) repository.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUser() error {
	// TODO: impl
	return nil
}
