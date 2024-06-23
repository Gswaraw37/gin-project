package config

import (
	appconfig "gin-project/config/app_config"
	dbconfig "gin-project/config/db_config"
)

func InitConfig() {
	appconfig.InitAppConfig()
	dbconfig.InitDBConfig()
}
