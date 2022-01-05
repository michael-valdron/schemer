package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeSchemer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
		})
	})
	router.Run()
}
