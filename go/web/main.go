package main

import (
	"github.com/gin-gonic/gin"
	"visible/console"
	"visible/controllers"
	"visible/models"
)

func main() {
	r := gin.Default()
	r.Static("/src", "./src")
	r.LoadHTMLGlob("./templates/**/*")
	r.LoadHTMLFiles("./templates/index.tmpl", "./templates/topic/addSubscribe.tmpl")

	var indexController controllers.IndexContrller
	var TopicController controllers.TopicController

	r.GET("/", indexController.Index)
	r.GET("/topic/index", TopicController.Index)
	r.GET("/topic/addSubscribe", TopicController.AddSubscribe)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := models.Init()
	if err != nil {
		panic(err)
	}
	defer models.Close()
	go func() {
		subscribe()
	}()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func subscribe() {
	console.Subscribe()
}

