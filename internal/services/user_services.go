package services

import (
	"errors"

	"github.com/sigit14ap/user-service/helpers"
	"github.com/sigit14ap/user-service/internal/domain"
	repository "github.com/sigit14ap/user-service/internal/repository/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email string, phone string, password string) error
	Login(emailOrPhone string, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	jwtSecret      string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) UserService {
	return &userService{
		userRepository: userRepo,
		jwtSecret:      jwtSecret,
	}
}

func (service *userService) Register(email, phone, password string) error {

	if email != "" {
		_, err := service.userRepository.GetUserByEmail(email)
		if err == nil {
			return errors.New("email already registered")
		}
	}

	if phone != "" {
		_, err := service.userRepository.GetUserByPhone(phone)
		if err == nil {
			return errors.New("phone number already registered")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email:    email,
		Phone:    phone,
		Password: string(hashedPassword),
	}

	return service.userRepository.CreateUser(user)
}

func (service *userService) Login(emailOrPhone, password string) (string, error) {
	var user *domain.User
	var err error

	if helpers.IsValidEmail(emailOrPhone) {
		user, err = service.userRepository.GetUserByEmail(emailOrPhone)
	} else {
		user, err = service.userRepository.GetUserByPhone(emailOrPhone)
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := helpers.GenerateJWT(user.Email, user.Phone)

	if err != nil {
		return "", err
	}

	return token, nil
}
