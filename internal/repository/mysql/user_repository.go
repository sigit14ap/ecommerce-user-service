package repository

import (
	"github.com/sigit14ap/user-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByPhone(phone string) (*domain.User, error)
	GetUserById(id uint64) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) CreateUser(user *domain.User) error {
	return repository.db.Create(user).Error
}

func (repository *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := repository.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (repository *userRepository) GetUserByPhone(phone string) (*domain.User, error) {
	var user domain.User
	err := repository.db.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func (repository *userRepository) GetUserById(id uint64) (*domain.User, error) {
	var shop domain.User
	err := repository.db.Where("id = ?", id).First(&shop).Error
	return &shop, err
}
