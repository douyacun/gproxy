package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Github *_Github

type _Github struct{}

type PostRequest struct {
	Url     string            `json:"url"`
	Body    string            `json:"body"`
	Header  map[string]string `json:"header"`
	Timeout int               `json:"timeout"`
}

func (*_Github) Post(ctx *gin.Context) {
	var b PostRequest
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	req, err := http.NewRequest("post", b.Url, strings.NewReader(b.Body))
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	for k, v := range b.Header {
		req.Header.Set(k, v)
	}
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Duration(b.Timeout),
	}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	res, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		ctx.JSON(resp.StatusCode, res)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
