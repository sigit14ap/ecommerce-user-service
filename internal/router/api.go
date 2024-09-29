package router

import (
	"github.com/gin-gonic/gin"
	delivery "github.com/sigit14ap/user-service/internal/delivery/http"
	"github.com/sigit14ap/user-service/internal/middleware"
)

func NewRouter(userHandler *delivery.UserHandler) *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1")
	v1.Use(middleware.ServiceMiddleware())

	users := v1.Group("/users")
	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)

		user := users.Use(middleware.AuthMiddleware())
		{
			user.GET("/me", userHandler.Me)
		}
	}

	return router
}
