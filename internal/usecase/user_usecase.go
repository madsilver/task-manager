package usecase

import (
	"context"
	"fmt"
	"github.com/madsilver/task-manager/internal/domain"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

// CreateUser creates a new user in the system.
// ctx: context for the operation
// user: pointer to the domain.User to be created
// Returns an error if the user input is invalid or if the creation fails.
func (uc *userUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return fmt.Errorf("invalid user input")
	}
	return uc.repo.Create(user)
}

// GetUserByEmail retrieves a user by their email address.
// ctx: context for the operation
// email: the email address of the user to retrieve
// Returns a pointer to the domain.User if found, or an error if not found or any other issue occurs.
func (uc *userUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uc.repo.GetByEmail(email)
}
