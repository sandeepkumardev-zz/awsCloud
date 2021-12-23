package controller

import (
	"awsCloud/database/models"
	"awsCloud/http/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sign In controller
// @Summary Sign In with credentials.
// @Description A registered user can sign in with their credentials.
// @Tags Sign In
// @Accept  json
// @Produce  json
// @Param user body models.SignInUser true "Sign In User"
// @Success 200 {object} models.SignInUser
// @Failure 401 {object} object
// @Router /signin [post]
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

// Sign Up controller
// @Summary Sign Up with credentials.
// @Description A new user can sign up with their username, password & phone number.
// @Tags Sign Up
// @Accept  json
// @Produce  json
// @Param user body models.User true "Sign Up User"
// @Success 200 {object} models.User
// @Failure 400 {object} object
// @Router /signup [post]
func SignUp(ctx *gin.Context) {
	var user models.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, models.Response{Success: false, Message: "Invalid input provided", Data: nil})
		return
	}

	res, status := services.CreateUser(&user)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

// Refresh token controller
// @Summary Varify token & create a new token.
// @Description You need to signedIn and give a Token in headers then "Refresh Token" will execute.
// @Tags Refresh token
// @Accept  json
// @Produce  json
// @Router /refreshToken [post]
func RefreshToken(ctx *gin.Context) {

}

// Varify OTP controller
// @Summary Varify OTP
// @Description You need to provide a Token in body.
// @Tags Varify OTP
// @Accept  json
// @Produce  json
// @Param user body models.OTP true "Varify OTP"
// @Success 200 {object} models.OTP
// @Failure 400 {object} object
// @Router /refreshToken [post]
func VerifyOTP(ctx *gin.Context) {
	var otp models.OTP

	if credErr := ctx.ShouldBindJSON(&otp); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, models.Response{Success: false, Message: "Invalid input provided", Data: nil})
		return
	}

	res, status := services.VerifyOTP(&otp)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}
