package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	router.POST("/github/post", Github.Post)
}
