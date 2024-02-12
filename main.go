package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	config "github.com/ValSpp/gin-gorm-restapi/configs"
)

func main() {
	config.InitialMigration()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	r.Run()
}
