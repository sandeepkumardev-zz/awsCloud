package controller

import (
	"awsCloud/database/models"
	"awsCloud/http/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	var user models.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, status := services.VerifyUser(&user)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

func SignUp(ctx *gin.Context) {
	var user models.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, status := services.CreateUser(&user)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

func RefreshToken(ctx *gin.Context) {

}
