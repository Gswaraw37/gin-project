package uploadcontroller

import (
	"fmt"
	appconfig "gin-project/config/app_config"
	"gin-project/utilities"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleUploadFile(ctx *gin.Context) {
	claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email => ", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userId => ", userId)

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "File is required",
		})

		return
	}

	// VALIDATION FILE BY EXTENTION
	fileExtention := []string{".jpg", ".jpeg", ".png", ".gif"}
	isValidated := utilities.FileValidationExtention(fileHeader, fileExtention)
	if !isValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "File type is not allowed",
		})

		return
	}

	// VALIDATION FILE BY CONTENT-TYPE
	// fileType := []string{"image/jpg", "image/jpeg", "image/png", "image/gif"}
	// isValidated := utilities.FileValidation(fileHeader, fileType)
	// if !isValidated {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "File type is not allowed",
	// 	})
	// }

	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := utilities.RandomFileName(extensionFile)

	isSaved := utilities.SaveFile(ctx, fileHeader, filename)
	if !isSaved {
		ctx.JSON(500, gin.H{
			"message": "Save File Failed",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success Upload File",
		"data":    filename,
	})
}

func HandleRemoveFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if filename == "" {
		ctx.JSON(400, gin.H{
			"message": "File name is required",
		})
	}

	err := utilities.RemoveFile(appconfig.DIR_FILE + filename)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "File Deleted",
	})
}
