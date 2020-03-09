package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopicController struct {

}

func (Topic *TopicController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "topic/index.tmpl", gin.H{})
}


func (Topic *TopicController) AddSubscribe(c *gin.Context) {
	c.HTML(http.StatusOK, "topic/addSubscribe.tmpl", gin.H{
		"title": "Main website",
	})
}

func (Topic *TopicController) GetTopicList(c *gin.Context) {

}