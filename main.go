package main

import (
	"myapp/db"
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	defer db.DB.Close() 

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
