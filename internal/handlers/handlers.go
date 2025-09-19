package handlers

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Home",
	})
}

func AboutHandler(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title": "Acerca de mí",
	})
}

func ContactHandler(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"title": "Contacto",
	})
}
