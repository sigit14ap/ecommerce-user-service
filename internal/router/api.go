package router

import (
	delivery "github.com/sigit14ap/user-service/internal/delivery/http"
	"github.com/sigit14ap/user-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *delivery.UserHandler, jwtSecret string) *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		v1.Use(middleware.AuthMiddleware(jwtSecret))
	}

	return router
}
