package router

import (
	"github.com/gin-gonic/gin"
	log2 "github.com/spigcoder/Hydra/service/internal/log"
)

type Group interface {
	AddRoutes(api *gin.RouterGroup)
}

type LogRouter struct {
	log *log2.Log
}

func NewLogRouter(log *log2.Log) Group {
	return &LogRouter{
		log: log,
	}
}

func (r *LogRouter) AddRoutes(api *gin.RouterGroup) {
	logHandler := log2.NewHTTPHandler(r.log)

	logApis := api.Group("/log")

	{
		logApis.GET("", logHandler.Consume)
		logApis.POST("", logHandler.Produce)
	}
}
