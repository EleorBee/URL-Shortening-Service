package endpoint

import (
	"URLShort/internal/model"
	"URLShort/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Endpoint struct {
	s *service.Service
}

func New(s *service.Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) CreateURL(ctx *gin.Context) {

	data, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var url Url

	err = json.Unmarshal(data, &url)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	surl, err := e.s.CreateUrl(url.Url)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, surl)
}

func (e *Endpoint) DeleteURL(ctx *gin.Context) {
	code := ctx.Param("url")

	err := e.s.DeleteUrl(code)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (e *Endpoint) GetURL(ctx *gin.Context) {
	code := ctx.Param("url")

	data, err := e.s.GetUrl(code)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
	ctx.Redirect(http.StatusMovedPermanently, data.Url)
	ctx.Abort()
}
func (e *Endpoint) GetURLStats(ctx *gin.Context) {
	code := ctx.Param("url")
	data := new(model.ShortUrl)
	data, err := e.s.GetUrlStats(code)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
func (e *Endpoint) UpdateURL(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	code := ctx.Param("url")

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var url Url

	err = json.Unmarshal(data, &url)

	surl, err := e.s.UpdateUrl(code, url.Url)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, surl)
}
