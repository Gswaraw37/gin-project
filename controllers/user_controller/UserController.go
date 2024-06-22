package usercontroller

import "github.com/gin-gonic/gin"

func GetAllUser(ctx *gin.Context) {
	isValidated := true

	if !isValidated {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Bad Request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"hello": "user",
	})
}
