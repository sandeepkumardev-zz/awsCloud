package services

import (
	"awsCloud/http/repositry"
	"awsCloud/http/utils"

	"github.com/gin-gonic/gin"
)

func CreateBucket(ctx *gin.Context) (res Response, status int) {
	username := ctx.Request.Header.Get("username")
	userId := ctx.Request.Header.Get("userId")

	bucketName := utils.CreateBucketName(username, userId)
	// createing new bucket
	_, BucketErr := repositry.CreateBucket(bucketName)
	if BucketErr != nil {
		return Response{Success: false, Message: BucketErr.Error(), Data: nil}, 400
	}

	return Response{Success: true, Message: "Bucket created successfully!", Data: nil}, 200
}

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

	if resp.Contents == nil {
		return Response{Success: true, Message: "Bucket is empty!", Data: resp.Contents}, 200
	}

	return Response{Success: true, Message: "Successfully fetched", Data: resp.Contents}, 200
}

// func GetObject(filename string) {
// 	fmt.Println("Downloading: ", filename)

// 	resp, err := s3session.GetObject(&s3.GetObjectInput{
// 		Bucket: aws.String(BUCKET_NAME),
// 		Key:    aws.String(filename),
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	err = ioutil.WriteFile(filename, body, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func DeleteObject(filename string) (resp *s3.DeleteObjectOutput) {
// 	fmt.Println("Deleting: ", filename)
// 	resp, err := s3session.DeleteObject(&s3.DeleteObjectInput{
// 		Bucket: aws.String(BUCKET_NAME),
// 		Key:    aws.String(filename),
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	return resp
// }
