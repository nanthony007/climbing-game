package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanthony007/climbing-game/pkg"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/grade", func(c *gin.Context) {
		input := pkg.Input{Grade: "v5", Seconds: 400.25}
		output, _ := input.Compute()
		c.JSON(200, gin.H{
			"output": output,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
