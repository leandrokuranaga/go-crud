package main

import (
	"log"
	_ "myapp/docs"
	database "myapp/internal/database"
	"myapp/internal/handlers/events"
	"myapp/internal/handlers/register"
	"myapp/internal/handlers/users"
	"myapp/internal/repository"
	"os"

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
	db, err := database.InitDb()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	eventRepo := repository.NewEventRepository(db)
	userRepo := repository.NewUserRepository(db)

	server := gin.Default()

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	events.RegisterRoutes(server, eventRepo)
	register.RegisterRoutes(server, eventRepo)
	users.RegisterRoutes(server, userRepo)

	port := os.Getenv("PORT")

	server.Run(port)
}
