package repository

import (
	"project/internal/entity"

	"gorm.io/gorm"
)

// UserRepository interface defines methods for user-related database operations
type UserRepository interface {
	Create(user *entity.User) error // Create a new user
	GetAll() ([]entity.User, error) // Get all users
}

// userRepository implements the UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user into the database
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// GetAll retrieves all users from the database
func (r *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}
