package services

import (
	"awsCloud/database/models"
	"awsCloud/http/repositry"
	"awsCloud/http/utils"
	"awsCloud/http/validator"
	verification "awsCloud/http/verification"
	"log"
	"os"
	"strings"

	uuid "github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Response models.Response

func VerifyUser(user *models.SignInUser) (res Response, status int) {
	vErr := validator.SignInValidator(user)
	if vErr != nil {
		return Response{Success: false, Message: vErr.Error(), Data: nil}, 400
	}

	item, err := repositry.GetItem(user.Username)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	if err = bcrypt.CompareHashAndPassword([]byte(item["password"].(string)), []byte(user.Password)); err != nil {
		return Response{Success: false, Message: "Wrong password!", Data: nil}, 401
	}

	token, err := utils.CreateToken(user.Username, item["id"].(string), item["verified"].(string))
	if err != nil {
		return Response{Message: "Something went wrong!", Data: nil, Success: false}, 200
	}

	tokens := map[string]string{
		"access_token":              token.AccessToken,
		"refresh_token":             token.RefreshToken,
		"access_token_expire_time":  os.Getenv("EXPIRE_ACCESS_TIME") + " minutes",
		"refresh_token_expire_time": os.Getenv("EXPIRE_REFRESH_TIME") + " minutes",
	}

	return Response{Success: true, Message: "SignIn successful!", Data: tokens}, 200
}

func CreateUser(user *models.User) (res Response, status int) {
	vErr := validator.SignUpValidator(user)
	if vErr != nil {
		return Response{Success: false, Message: vErr.Error(), Data: nil}, 400
	}

	// check if user exists
	exists, err := repositry.FindItem(user.Username)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 200
	}
	if exists {
		return Response{Success: false, Message: "User already exists!", Data: nil}, 200
	}

	// generate new user id
	newId, _ := uuid.NewV4()
	// hash password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if hashErr != nil {
		return Response{Success: false, Message: "Something went wrong!", Data: nil}, 200
	}

	// assign id & hash password
	user.Id = newId.String()
	user.Password = string(hash)

	firstName := strings.Split(user.Username, " ")[0]
	bucketName := utils.CreateBucketName(firstName, user.Id)
	// createing new bucket
	_, BucketErr := repositry.CreateBucket(bucketName)
	if BucketErr != nil {
		return Response{Success: false, Message: BucketErr.Error(), Data: nil}, 400
	}

	// Add new user in database
	err = repositry.PutItem(user)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	otp, err := utils.CreateOTP(user.Id)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	log.Println("OTP - ", otp)

	// isSend := verification.SendSMS(user.PhoneNumber, "Welcome to awsCloud services. OTP - "+strconv.Itoa(otp))
	// if !isSend {
	// 	log.Println("Message failed to send - ", user.PhoneNumber)
	// }

	return Response{Success: true, Message: "SignUp successful!", Data: user}, 200
}

func VerifyOTP(otp *models.OTP) (res Response, status int) {
	err := verification.VerifyOTP(otp.Id, otp.OTP)
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	dberr := repositry.UpadteItem(otp.Id, "true")
	if dberr != nil {
		return Response{Success: false, Message: dberr.Error(), Data: nil}, 400
	}

	return Response{Success: true, Message: "OTP varification successfully!", Data: nil}, 200
}
