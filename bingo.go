package bingo

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xylong/bingo/ioc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

type Bean interface {
	Name() string
}

// Bingo 脚手架
type Bingo struct {
	*gin.Engine
	group *gin.RouterGroup       // 路由分组
	expr  map[string]interface{} // 表达式
}

// Init 初始化
func Init() *Bingo {
	b := &Bingo{
		Engine: gin.New(),
		expr:   make(map[string]interface{}),
	}

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
			b.joinBean(c) // 将控制器加入容器
		}

		return b
	}
}

// Inject 注入依赖实体
func (b *Bingo) Inject(entities ...interface{}) *Bingo {
	ioc.Factory.Unwrap(entities...)
	return b
}

// applyBean 给ioc容器中所有结构体注入依赖
func (b *Bingo) applyBean() {
	for key, value := range ioc.Factory.GetMapper() {
		if key.Elem().Kind() == reflect.Struct {
			ioc.Factory.Apply(value.Interface())
		}
	}
}

// joinBean 加入bean
func (b *Bingo) joinBean(beans ...Bean) *Bingo {
	for _, bean := range beans {
		b.expr[bean.Name()] = bean
		ioc.Factory.Set(bean)
	}

	return b
}

// Crontab 定时任务
// cron 定时任务表达式
// expr 执行函数，执行函数如果是方法则直接加入定时任务方法，如果是gin表达式则转为可执行方法加入定时任务
func (b *Bingo) Crontab(cron string, expr interface{}) *Bingo {
	var err error

	switch value := expr.(type) {
	case func():
		_, err = getCron().AddFunc(cron, value)
	case string:
		_, err = getCron().AddFunc(cron, func() {
			if _, err := ExecExpr(Expr(value), b.expr); err != nil {
				logrus.Error(err.Error())
			}
		})
	}

	if err != nil {
		logrus.Error(err.Error())
	}

	return b
}

// Lunch 启动
func (b *Bingo) Lunch() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: b.Engine,
	}

	b.applyBean()

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
