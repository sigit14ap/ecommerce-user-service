package router

import (
	"github.com/gin-gonic/gin"
	delivery "github.com/sigit14ap/user-service/internal/delivery/http"
	"github.com/sigit14ap/user-service/internal/middleware"
)

func NewRouter(userHandler *delivery.UserHandler) *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		_ = v1.Use(middleware.AuthMiddleware())
		{

		}
	}

	return router
}
