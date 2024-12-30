package usecase

import (
	"project/internal/entity"
	"project/internal/repository"
)

// UserUsecase interface defines methods for user-related business logic
type UserUsecase interface {
	CreateUser(user *entity.User) error  // Create a new user
	GetAllUsers() ([]entity.User, error) // Get all users
}

// userUsecase implements the UserUsecase interface
type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase creates a new instance of userUsecase
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

// CreateUser handles the business logic for creating a new user
func (u *userUsecase) CreateUser(user *entity.User) error {
	return u.userRepo.Create(user)
}

// GetAllUsers retrieves all users via the repository
func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
	return u.userRepo.GetAll()
}
