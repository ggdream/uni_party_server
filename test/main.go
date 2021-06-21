package main

import (
	"github.com/gin-gonic/gin"
)


type Greet struct {
	Name	string	`json:"name" form:"name"`
	Age		int8	`json:"age" form:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var greet Greet
		if err := c.ShouldBind(&greet); err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"name": greet.Name,
			"age": greet.Age,
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
