package usecase

import (
	"errors"

	"github.com/sigit14ap/user-service/helpers"
	"github.com/sigit14ap/user-service/internal/domain"
	repository "github.com/sigit14ap/user-service/internal/repository/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(email string, phone string, password string) error
	Login(emailOrPhone string, password string) (string, error)
	Me(id uint64) (*domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (uc *userUsecase) Register(email, phone, password string) error {

	if email != "" {
		_, err := uc.userRepository.GetUserByEmail(email)
		if err == nil {
			return errors.New("email already registered")
		}
	}

	if phone != "" {
		_, err := uc.userRepository.GetUserByPhone(phone)
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

	return uc.userRepository.CreateUser(user)
}

func (uc *userUsecase) Login(emailOrPhone, password string) (string, error) {
	var user *domain.User
	var err error

	if helpers.IsValidEmail(emailOrPhone) {
		user, err = uc.userRepository.GetUserByEmail(emailOrPhone)
	} else {
		user, err = uc.userRepository.GetUserByPhone(emailOrPhone)
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := helpers.GenerateJWT(user.Email, user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUsecase) Me(id uint64) (*domain.User, error) {
	user, err := uc.userRepository.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
