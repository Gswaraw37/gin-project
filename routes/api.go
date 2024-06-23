package routes

import (
	authcontroller "gin-project/App/controllers/auth_controller"
	bookcontroller "gin-project/App/controllers/book_controller"
	uploadcontroller "gin-project/App/controllers/upload_controller"
	usercontroller "gin-project/App/controllers/user_controller"
	appconfig "gin-project/config/app_config"
	"gin-project/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)

	// AUTH ROUTE
	route.POST("/login", authcontroller.Login)

	// ROUTE USER
	userRoute := route.Group("/user")
	userRoute.GET("/", usercontroller.Index)
	userRoute.GET("/paginate", usercontroller.IndexPaginate)
	userRoute.GET("/:id", usercontroller.Show)
	userRoute.POST("/", usercontroller.Store)
	userRoute.PUT("/:id", usercontroller.Update)
	userRoute.DELETE("/:id", usercontroller.Delete)

	// ROUTE BOOK
	route.GET("/book", bookcontroller.GetAllBook)

	// ROUTE FILE
	authRoute := route.Group("/file", middleware.AuthMiddleware)
	authRoute.POST("/", uploadcontroller.HandleUploadFile)
	authRoute.DELETE("/:filename", middleware.AuthMiddleware, uploadcontroller.HandleRemoveFile)
}
