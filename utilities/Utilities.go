package utilities

import (
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	rand.Seed(time.Now().UnixMilli())
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}

	return result
}

func FileValidationExtention(fileHeader *multipart.FileHeader, fileExtention []string) bool {
	extention := filepath.Ext(fileHeader.Filename)
	result := false

	for _, typeFile := range fileExtention {
		if extention == typeFile {
			result = true
			break
		}
	}

	return result
}

func RandomFileName(extensionFile string, prefix ...string) string {
	currentPrefix := "file"
	if len(prefix) > 0 {
		if prefix[0] != "" {
			currentPrefix = prefix[0]
		}
	}

	filename := fmt.Sprintf("%s-%s%s", currentPrefix, RandomString(5), extensionFile)

	return filename
}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {
	err := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))
	if err != nil {
		log.Println("File Upload Failed")

		return false
	} else {
		return true
	}
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("Remove File Failed")

		return err
	}

	return nil
}
