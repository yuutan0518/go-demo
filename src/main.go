package main

import (
	"github.com/yuutan0518/go-demo/db"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	db.Init()
	r.Run()

	db.Close()

}
