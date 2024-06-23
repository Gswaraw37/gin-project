package routes

import (
	bookcontroller "gin-project/App/controllers/book_controller"
	uploadcontroller "gin-project/App/controllers/upload_controller"
	usercontroller "gin-project/App/controllers/user_controller"
	appconfig "gin-project/config/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)

	// ROUTE USER
	route.GET("/user", usercontroller.Index)
	route.GET("/user/paginate", usercontroller.IndexPaginate)
	route.GET("/user/:id", usercontroller.Show)
	route.POST("/user", usercontroller.Store)
	route.PUT("/user/:id", usercontroller.Update)
	route.DELETE("/user/:id", usercontroller.Delete)

	// ROUTE BOOK
	route.GET("/book", bookcontroller.GetAllBook)

	// ROUTE FILE
	route.POST("/file", uploadcontroller.HandleUploadFile)
	route.DELETE("/file/:name", uploadcontroller.HandleRemoveFile)
}
