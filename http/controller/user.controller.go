package controller

import (
	"awsCloud/database/models"
	"awsCloud/http/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	var user models.SignInUser

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, models.Response{Success: false, Message: "Invalid input provided", Data: nil})
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
		ctx.JSON(http.StatusUnprocessableEntity, models.Response{Success: false, Message: "Invalid input provided", Data: nil})
		return
	}

	fmt.Println(user)

	res, status := services.CreateUser(&user)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

func RefreshToken(ctx *gin.Context) {

}
