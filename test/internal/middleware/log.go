package middleware

import (
	"fmt"
	"github.com/xylong/bingo"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Before(ctx *bingo.Context) error {
	fmt.Println("假装写日志")
	return nil
}