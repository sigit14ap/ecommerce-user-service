package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	helpers "github.com/sigit14ap/user-service/helpers"
	"github.com/sigit14ap/user-service/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
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

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
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

	err = handler.userUsecase.Register(registerRequest.Email, registerRequest.Phone, registerRequest.Password)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
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

	token, err := handler.userUsecase.Login(loginRequest.EmailOrPhone, loginRequest.Password)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Email/Phone number or password are incorrect")
		return
	}

	helpers.SuccessResponse(context, gin.H{"token": token})
}

func (handler *UserHandler) Me(context *gin.Context) {
	userID, exists := context.Get("userID")

	if !exists {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "User id not found")
		return
	}

	userIDUint, valid := userID.(uint64)
	if !valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is not valid"})
		return
	}

	user, err := handler.userUsecase.Me(userIDUint)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "User data not found")
		return
	}

	helpers.SuccessResponse(context, gin.H{"user": user})
}
