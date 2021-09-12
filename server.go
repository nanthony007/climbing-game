package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanthony007/climbing-game/pkg/models"
	"log"
)

// This is a pattern for API use
//oldPoints := i.TotalPoints()
//i.Routes = append(i.Routes, route)
//newPoints := i.TotalPoints()

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/problem", func(c *gin.Context) {
		var input models.Input
		// If `GET`, only `Form` binding engine (`query`) used.
		// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
		// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
		err := c.ShouldBind(&input)
		if err != nil {
			log.Fatal("Could not assemble data into json.")
		}
		newData := input.Compute()
		fmt.Println(newData)
		c.JSON(201, gin.H{
			"output": newData,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
