package bootstrap

import (
	"gin-project/config"
	appconfig "gin-project/config/app_config"
	"gin-project/database"
	"gin-project/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	// LOAD .ENV FILE
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// INIT CONFIG
	config.InitConfig()

	// DATABASE CONNECTION
	database.ConnectDatabase()
	// INIT GIN ENGINE
	app := gin.Default()

	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(":" + appconfig.PORT)
}
