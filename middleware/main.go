package middleware

import (
	"awsCloud/models"
	"awsCloud/utils"

	"github.com/gin-gonic/gin"
)

func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//verify Token
		_, err := utils.VerifyAccessToken(c)
		if err != "" {
			c.JSON(401, models.Response{Success: true, Message: err, Data: nil})
			c.Abort()
		}
		c.Next()
	}
}
