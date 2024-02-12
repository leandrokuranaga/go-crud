package main

import (
	"myapp/db"
	_ "myapp/docs"
	"myapp/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CRUD User API
// @version 1.0
// @description A CRUD User API in go using gin framework

// @host localhost:8080
// @BasePath /api
func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	defer db.DB.Close()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
