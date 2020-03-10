package controllers

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"gproxy/internal/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Proxy *_proxy

type _proxy struct{}

// Data contains ContentType and bytes data.
type Data struct {
	ContentType string
	Data        []byte
}

type PostRequest struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Body    string            `json:"body"`
	Header  map[string]string `json:"header"`
	Timeout int               `json:"timeout"`
}

// Render (Data) writes data with custom ContentType.
func (d Data) Render(w http.ResponseWriter) (err error) {
	d.WriteContentType(w)
	_, err = w.Write(d.Data)
	return
}

// WriteContentType (Data) writes custom ContentType.
func (d Data) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{d.ContentType}
	}
}

func (*_proxy) File(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.String(http.StatusBadRequest, "url参数缺失")
		return
	}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	//ctx.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Render(resp.StatusCode, Data{
		ContentType: resp.Header.Get("Content-Type"),
		Data:        body,
	})
	return
}

func (*_proxy) Request(ctx *gin.Context) {
	var b PostRequest
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	req, err := http.NewRequest(b.Method, b.Url, strings.NewReader(b.Body))
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	logger.Infof("request: \nmethod: %s\nurl: %s\nheader: %v\nbody: %s", b.Method, b.Url, b.Header, b.Body)
	for k, v := range b.Header {
		req.Header.Set(k, v)
	}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(b.Timeout),
	}
	resp, err := client.Do(req)
	if err != nil {
		ctx.String(500, err.Error())
		return
	}
	res, _ := ioutil.ReadAll(resp.Body)
	ctx.Render(resp.StatusCode, Data{
		ContentType: resp.Header.Get("Content-Type"),
		Data:        res,
	})
	return
}
