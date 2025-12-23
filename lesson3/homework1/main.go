package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/talk", func(c *gin.Context) {
		str := c.Query("msg")
		switch str {
		case "ping":
			c.JSON(200, gin.H{
				"data": "pong",
			})
		case "helloserver":
			c.JSON(200, gin.H{
				"data": "helloclient",
			})
		default:
			c.JSON(200, gin.H{
				"data": "unknown message",
			})
		}
	})
	r.Run(":80")
}
