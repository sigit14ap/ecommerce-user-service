package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	helpers "github.com/sigit14ap/user-service/helpers"
	"github.com/sigit14ap/user-service/internal/services"
)

type UserHandler struct {
	userService services.UserService
}

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone" validate:"required"`
	Password     string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

var validate *validator.Validate

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) Register(context *gin.Context) {
	validate = validator.New()
	var registerRequest RegisterRequest
	err := context.BindJSON(&registerRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(registerRequest)
	if err != nil {
		helpers.ErrorValidationResponse(context, err)
		return
	}

	err = handler.userService.Register(registerRequest.Email, registerRequest.Phone, registerRequest.Password)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(context, gin.H{"message": "User created successfully"})
}

func (handler *UserHandler) Login(context *gin.Context) {
	validate = validator.New()
	var loginRequest LoginRequest
	err := context.BindJSON(&loginRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(loginRequest)
	if err != nil {
		helpers.ErrorValidationResponse(context, err)
		return
	}

	token, err := handler.userService.Login(loginRequest.EmailOrPhone, loginRequest.Password)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Email/Phone number or password are incorrect")
		return
	}

	helpers.SuccessResponse(context, gin.H{"token": token})
}
