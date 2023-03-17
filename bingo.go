package bingo

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo/iface"
	"github.com/xylong/bingo/ioc"
	"go.uber.org/zap"
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
func (b *Bingo) Mount(group string, controller ...Controller) func(middleware ...iface.Middleware) *Bingo {
	b.group = b.Group(group)
	g := NewGroup(b.group)

	return func(middleware ...iface.Middleware) *Bingo {
		for _, c := range controller {
			g.attach(middleware...)
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
				zap.L().Error("crontab express error", zap.Error(err))
			}
		})
	}

	if err != nil {
		zap.L().Error("crontab error", zap.Error(err))
	}

	return b
}

// Lunch 启动
func (b *Bingo) Lunch(port int) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: b.Engine,
	}

	b.applyBean()

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorf("http listen %d: %s\n", port, err)
		}
	}()

	getCron().Start()

	// 等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server forced to shutdown:", zap.Error(err))
	}

	fmt.Println("server exiting")
}
