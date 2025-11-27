package main

import "github.com/xiaohangshuhub/gin"

func main() {

	e := gin.New()
	e.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	e.Run()
}
