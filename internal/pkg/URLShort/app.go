package URLShort

import (
	"URLShort/internal/endpoint"
	"URLShort/internal/service"
	"URLShort/internal/storage/sqlite"
	"github.com/gin-gonic/gin"
)

type App struct {
	endpoint *endpoint.Endpoint
	router   *gin.Engine
}

func New() *App {

	route := gin.Default()

	strg, err := sqlite.New("storage/Storage.db")

	if err != nil {
		panic(err)
	}

	srv := service.New(strg)

	end := endpoint.New(srv)

	route.POST("/shorten", end.CreateURL)
	route.PUT("/shorten/:url", end.UpdateURL)
	route.GET("/shorten/:url", end.GetURL)
	route.GET("/shorten/:url/stats", end.GetURLStats)
	route.DELETE("/shorten/:url", end.DeleteURL)

	return &App{
		router:   route,
		endpoint: end,
	}
}

func (a *App) Run() {
	err := a.router.Run(":5050")

	if err != nil {
		panic(err)
	}
}
