package bingo

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Bingo 脚手架
type Bingo struct {
	*gin.Engine
	group *gin.RouterGroup
}

// Init 初始化
func Init() *Bingo {
	b := &Bingo{Engine: gin.New()}
	return b
}

// Route 路由
func (b *Bingo) Route(group string, callback func(group *Group)) *Bingo {
	callback(NewGroup(b.Engine.Group(group)))
	return b
}

// Mount 挂载控制器
func (b *Bingo) Mount(group string, controller ...Controller) func(middleware ...Middleware) *Bingo {
	b.group = b.Group(group)
	g := NewGroup(b.group)

	return func(middleware ...Middleware) *Bingo {
		for _, c := range controller {
			g.middlewares = append(g.middlewares, middleware...)
			c.Route(g)
		}

		return b
	}
}

// Crontab 定时任务
// 0/1 * * * * * 每秒执行
func (b *Bingo) Crontab(expression string, f func()) *Bingo {
	if entryID, err := getCron().AddFunc(expression, f); err != nil {
		logrus.Error(err.Error())
	} else {
		logrus.Info("cron id is ", entryID)
	}

	return b
}

// Lunch 启动
func (b *Bingo) Lunch() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: b.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Info("http server closed")
		}
	}()

	getCron().Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Info("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
