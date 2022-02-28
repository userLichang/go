package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "我无敌了")
	})
	r.Run()
}
