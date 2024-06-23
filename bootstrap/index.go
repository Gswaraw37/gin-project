package bootstrap

import (
	"gin-project/config"
	appconfig "gin-project/config/app_config"
	corsconfig "gin-project/config/cors_config"
	logconfig "gin-project/config/log_config"
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
	// LOGGING
	logconfig.DefaultLogging()
	// INIT GIN ENGINE
	app := gin.Default()

	// CORS PACKAGE
	app.Use(corsconfig.CorsConfigContrib())
	// CORS MANUAL
	// app.Use(corsconfig.CorsConfig)

	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(":" + appconfig.PORT)
}
