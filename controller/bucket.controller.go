package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UploadItem(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Upload")
}
