package routes

import (
	bookcontroller "gin-project/App/controllers/book_controller"
	usercontroller "gin-project/App/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	// ROUTE USER
	route.GET("/user", usercontroller.Index)
	route.GET("/user/:id", usercontroller.Show)
	route.POST("/user", usercontroller.Store)
	route.PUT("/user/:id", usercontroller.Update)
	route.DELETE("/user/:id", usercontroller.Delete)

	// ROUTE BOOK
	route.GET("/book", bookcontroller.GetAllBook)
}
