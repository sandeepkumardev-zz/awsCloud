package services

import (
	"awsCloud/http/repositry"
	"awsCloud/http/utils"
	"strconv"

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
	verified := ctx.Request.Header.Get("verified")

	boolValue, _ := strconv.ParseBool(verified)
	if !boolValue {
		return Response{Success: false, Message: "You are not authorized user.", Data: nil}, 400
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return Response{Success: false, Message: err.Error(), Data: nil}, 415
	}
	files := form.File["myFile"]

	for _, file := range files {
		f, fErr := file.Open()
		if fErr != nil {
			return Response{Success: false, Message: err.Error(), Data: nil}, 415
		}
		defer f.Close()

		bucketName := utils.CreateBucketName(username, userId)
		err := repositry.UploadObject(bucketName, f, file.Filename)

		if err != nil {
			return Response{Success: false, Message: err.Error(), Data: nil}, 400
		}

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
