package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexContrller struct {

}

func (index *IndexContrller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}