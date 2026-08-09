package main

import (
	"github.com/adamwahyuh/banking/config"
	"github.com/adamwahyuh/banking/database"
	"github.com/adamwahyuh/banking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectDb()

	route := gin.Default()

	routes.RegisterApiRoute(route)
	route.Run(config.Getenv("APP_HOST") + ":" + config.Getenv("APP_PORT"))
}
