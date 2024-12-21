package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nandonyata/kerjaaja/common"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	common.Load(".")
	if common.Get().General.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := common.ConnectDB(logrus.New()); err != nil {
		logger.Fatalf("error connecting database. error: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := common.Get().General.Port
	if err := r.Run(port); err != nil {
		log.Fatalf("error listening on port: %v, error: %v", port, err)
	}
}
