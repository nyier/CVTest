package routes

import (
	"cvTest/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/about", handlers.AboutHandler)
	r.GET("/contact", handlers.ContactHandler)
	r.GET("/", handlers.IndexHandler)
}
