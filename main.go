package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, %s!", "joe")
	})
	fmt.Println(errors.New("text string"))
	r.Run(":8080")
}
