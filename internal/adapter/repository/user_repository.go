package repository

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/madsilver/task-manager/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type users struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Roles    Roles `gorm:"type:json" json:"roles"`
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

// Create inserts a new user into the database.
// user: the domain.User to be created
// Returns an error if the operation fails.
func (r *userRepository) Create(user *domain.User) error {
	model := users{Name: user.Name, Email: user.Email}
	if err := r.db.Create(&model).Error; err != nil {
		return err
	}
	user.ID = model.ID
	return nil
}

// GetByEmail retrieves a user by their email address from the database.
// email: the email address of the user to retrieve
// Returns a pointer to the domain.User if found, or an error if not found or any other issue occurs.
func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var model users
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return toDomainUser(model), nil
}

// toDomainUser converts a users model to a domain.User.
// model: the users model to convert
// Returns a pointer to the domain.User.
func toDomainUser(model users) *domain.User {
	return &domain.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Roles:    model.Roles,
		Password: model.Password,
	}
}

type Roles []string

func (s *Roles) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal roles: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

func (s *Roles) Value() (driver.Value, error) {
	return json.Marshal(s)
}
