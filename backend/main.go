package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nandonyata/kerjaaja/config"
)

func main() {
	// logger := logrus.New()

	config.Load(".")
	if config.Get().General.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := config.Get().General.Port
	if err := r.Run(port); err != nil {
		log.Fatalf("error listening on port: %v, error: %v", port, err)
	}
}
