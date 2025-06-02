package main

import (
	"first-project/config"
	"first-project/router"
)

func main() {
	config.InitConfig()
	router := router.SetUpRouter()

	router.Run(config.AppConfig.App.Port)
}
