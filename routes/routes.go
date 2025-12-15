package routes

import (
	"github.com/gin-gonic/gin"

	"student-api/handler"
	"student-api/middleware"
)

func RegisterRoutes(r *gin.Engine, h *handler.StudentHandler) {
	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Protected routes
	api := r.Group("/")
	api.Use(middleware.APIKeyMiddleware())

	{
		api.POST("/students", h.CreateStudent)
		api.GET("/students", h.GetStudents)
		api.PUT("/students/:id", h.UpdateStudent)
		api.DELETE("/students/:id", h.DeleteStudent)
	}
}
