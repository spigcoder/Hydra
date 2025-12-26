package main

import (
	"context"
	"errors"
	"github.com/spigcoder/Hydra/service/internal/log"
	"github.com/spigcoder/Hydra/service/internal/router"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/spigcoder/Hydra/pkg/config"
	"github.com/spigcoder/Hydra/pkg/logger"
)

func main() {
	if err := config.Init("configs", "config", "HYDRA"); err != nil {
		panic(err)
	}

	logger.Init("info")

	env := router.Env{
		Log: log.NewLog(),
	}

	r := router.NewRouter(env)
	addr := viper.GetString("addr")

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("ListenAndServe Failed", "error", err)
			panic(err)
		}
	}()

	// 监听推出信号，优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown failed", "error", err)
	}
	slog.Info("Server exiting")
}
