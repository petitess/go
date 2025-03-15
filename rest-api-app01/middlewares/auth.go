package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"site.org/abc/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token missing."})
		return
	}

	userId, err := utils.VarifiyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not create event. " + err.Error()})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
