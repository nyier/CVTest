package main

import (
	"cvTest/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("internal/templates/*")
	r.Static("/static", "./static")
	routes.RegisterRoutes(r)
	r.Run(":8000")

}
