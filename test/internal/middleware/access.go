package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
	"go.uber.org/zap"
	"time"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Before(ctx *bingo.Context) error {
	start := time.Now()
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery
	//form := ctx.Request.Form

	defer func() {
		zap.L().Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			//zap.Array("form", form),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", time.Since(start)),
		)
	}()

	ctx.Next()
	return nil
}
