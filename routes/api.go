package routes

import (
	bookcontroller "gin-project/controllers/book_controller"
	usercontroller "gin-project/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.GET("/user", usercontroller.GetAllUser)
	route.GET("/book", bookcontroller.GetAllBook)
}
