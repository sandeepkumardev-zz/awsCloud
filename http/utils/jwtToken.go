package utils

import (
	"awsCloud/database/models"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var td = &models.TokenDetails{}
var ACCESS_SECRET = os.Getenv("ACCESS_SECRET")
var REFRESH_SECRET = os.Getenv("REFRESH_SECRET")

func CreateToken(username string, id string, verified string) (*models.TokenDetails, error) {
	exAccTime, _ := strconv.Atoi(os.Getenv("EXPIRE_ACCESS_TIME"))
	exRefTime, _ := strconv.Atoi(os.Getenv("EXPIRE_REFRESH_TIME"))
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(exAccTime)).Unix()
	td.RtExpires = time.Now().Add(time.Minute * time.Duration(exRefTime)).Unix()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["expiresAt"] = td.AtExpires
	atClaims["username"] = username
	atClaims["userId"] = id
	atClaims["authorized"] = verified
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["expiresAt"] = td.RtExpires
	rtClaims["username"] = username
	atClaims["userId"] = id
	rtClaims["authorized"] = verified
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(REFRESH_SECRET))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func verifyToken(data string, secret []byte) (*jwt.Token, string) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, "Invalid Token!"
	}

	return token, ""
}

func VerifyAccessToken(ctx *gin.Context) (jwt.Claims, string) {
	if data := ctx.Request.Header.Get("Authorization"); data != "" {
		token, err := verifyToken(data, []byte(ACCESS_SECRET))
		if err != "" {
			return nil, err
		}

		// extract expiresTime from token
		ext := token.Claims.(jwt.MapClaims)
		expiresTime := ext["expiresAt"]
		username := ext["username"]
		userId := ext["userId"]
		verified := ext["authorized"]

		ctx.Request.Header.Set("username", fmt.Sprintf("%v", username))
		ctx.Request.Header.Set("userId", fmt.Sprintf("%v", userId))
		ctx.Request.Header.Set("verified", fmt.Sprintf("%v", verified))

		if int64(expiresTime.(float64)) < time.Now().Unix() {
			return nil, "Token expired!"
		}
		return token.Claims, ""
	} else {
		return nil, "You are not logged In"
	}
}
