package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test-task/pkg/sth"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(sth.CheckPingHeader())

	router.GET("/when/:year", sth.HandleYear)

	err := router.Run()
	if err != nil {
		log.Println("running error: ", err)
	}
}
