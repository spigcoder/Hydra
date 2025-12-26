package logger

import (
	"log/slog"
	"os"
	"strings"
)

// Init 初始化全局 JSON Logger
// levelStr: 日志级别 (debug, info, warn, error)
func Init(levelStr string) {
	// 1. 解析日志级别 (默认为 Info)
	var level slog.Level
	switch strings.ToLower(levelStr) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// 2. 配置 Handler 选项
	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	finalHandler := NewTraceHandler(handler)

	logger := slog.New(finalHandler)
	slog.SetDefault(logger)
}

type TraceHandler struct {
	slog.Handler
}

// NewTraceHandler 创建带 Trace 能力的 Handler
func NewTraceHandler(h slog.Handler) *TraceHandler {
	return &TraceHandler{Handler: h}
}
