package main

import (
	"log"

	_ "student-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"student-api/config"
	"student-api/handler"
	"student-api/repository"
	"student-api/routes"
	"student-api/service"
)

func main() {
	// Connect to PostgreSQL
	db := config.ConnectDB()
	defer db.Close()

	// Dependency Injection
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// Gin engine
	r := gin.New()

	// Register routes & middleware
	routes.RegisterRoutes(r, studentHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
