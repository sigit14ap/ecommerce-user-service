package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/user-service/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			helpers.ErrorResponse(context, http.StatusUnauthorized, "Unauthorized")
			context.Abort()
			return
		}

		_, err := helpers.ParseJWT(token)

		if err != nil {
			helpers.ErrorResponse(context, http.StatusUnauthorized, "Unauthorized")
			context.Abort()
			return
		}

		context.Next()
	}
}
