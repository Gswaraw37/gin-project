package bootstrap

import (
	appconfig "gin-project/config/app_config"
	"gin-project/database"
	"gin-project/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(appconfig.PORT)
}
