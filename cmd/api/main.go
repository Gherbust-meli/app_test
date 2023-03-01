package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	port := ":8001"
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"host_name": hostName,
			"port":      port,
		})
	})
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
