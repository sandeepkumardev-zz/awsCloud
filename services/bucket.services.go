package services

import (
	"awsCloud/repositry"
	"awsCloud/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UploadItem(ctx *gin.Context) (res Response, status int) {
	username := ctx.Request.Header.Get("username")
	userId := ctx.Request.Header.Get("userId")

	file, handler, fileErr := ctx.Request.FormFile("myFile")
	if fileErr != nil {
		return Response{Success: false, Message: fileErr.Error(), Data: nil}, 415
	}
	defer file.Close()

	bucketName := utils.CreateBucketName(username, userId)
	err := repositry.UploadObject(bucketName, file, handler.Filename)

	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	return Response{Success: true, Message: "File upload successfully!", Data: nil}, 200
}

func GetAllItem(ctx *gin.Context) (res Response, status int) {
	username := ctx.Request.Header.Get("username")
	userId := ctx.Request.Header.Get("userId")

	bucketName := utils.CreateBucketName(username, userId)
	resp, err := repositry.GetAllObjects(bucketName)

	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 400
	}

	fmt.Println(resp)
	return Response{Success: true, Message: "Successfully fetched", Data: resp}, 200
}
