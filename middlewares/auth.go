package middlewares

import (
	"net/http"

	"example.com/go_rest_api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Unauthorized"},
		)
		return
	}

	id, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Unauthorized"},
		)
		return
	}

	ctx.Set("id", id)

	ctx.Next()
}
