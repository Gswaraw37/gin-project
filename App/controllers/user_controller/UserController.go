package usercontroller

import (
	"gin-project/App/models"
	"gin-project/database"
	"gin-project/requests"
	"gin-project/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Data Tidak Ditemukan",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func Show(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(responses.UserResponse)

	// PAKE TABLE KARENA AGAR SPESIFIK MENGAMBIL DATA DARI TABLE MANA
	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Tidak Ditemukan",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	if err := ctx.ShouldBind(&userReq); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)

	if userEmailExist.Email != nil {
		ctx.JSON(400, gin.H{
			"message": "Email Already Exist",
		})

		return
	}

	user := new(models.User)
	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	err := database.DB.Table("users").Create(&user).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Create Data Failed",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    user,
	})
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)
	userEmailExist := new(models.User)

	if err := ctx.ShouldBind(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data Not Found",
		})

		return
	}

	// EMAIL EXISTS
	errUserEmail := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmail != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(400, gin.H{
			"message": "Email Already Exist",
		})

		return
	}

	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "Update Data Failed",
		})

		return
	}

	userResponse := responses.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    userResponse,
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)

	database.DB.Table("users").Where("id = ?", id).First(&user)
	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data Not Found",
		})

		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Delete Data Failed",
			"error":   errDb.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "Data Deleted",
		"data":    user,
	})
}

func IndexPaginate(ctx *gin.Context) {
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("perPage")
	if perPage == "" {
		perPage = "10"
	}

	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)

	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Data Tidak Ditemukan",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"data":    users,
		"page":    pageInt,
		"perPage": perPageInt,
	})
}
