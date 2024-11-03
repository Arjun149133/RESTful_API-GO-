package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "Hello Peter!",
		})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	r.GET("/pong", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "ping",
		})
	})
	r.Run()
}
