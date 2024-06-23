package authcontroller

import (
	"gin-project/App/models"
	"gin-project/database"
	"gin-project/requests"
	"gin-project/utilities"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {
	loginReq := new(requests.LoginRequest)

	if err := ctx.ShouldBind(&loginReq); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	user := new(models.User)
	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user).Error
	if errUser != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "credential is invalid",
		})

		return
	}

	// CHECK PASSWORD
	if loginReq.Password != "123456" {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "credential is invalid",
		})

		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utilities.GenerateToken(&claims)
	if errToken != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed generate token",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success login",
		"token":   token,
	})
}
