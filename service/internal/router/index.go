package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spigcoder/Hydra/service/internal/log"

	"github.com/spigcoder/Hydra/pkg/middleware"
)

type Env struct {
	Log *log.Log
}

func NewRouter(env Env) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ErrorHandler)
	api := r.Group("/" + viper.GetString("api.prefix"))

	logRouter := NewLogRouter(env.Log)
	logRouter.AddRoutes(api)

	return r
}
