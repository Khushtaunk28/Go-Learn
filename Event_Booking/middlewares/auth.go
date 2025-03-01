package middlewares

import (
	"Event_Booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticaton(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "jwt token empty"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Not authorized"})
		return
	}

	context.Set("userId",userId)

	context.Next()
}
