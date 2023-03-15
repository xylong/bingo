package iface

import "go.uber.org/zap/zapcore"

// Configure 配置
type Configure interface {
	LogConfigure

	GetMode() string
	GetPort() int
}

type LogConfigure interface {
	// Encoder 配置日志格式
	Encoder() (zapcore.EncoderConfig, bool)

	// Writer 配置日志输出
	Writer() (filename, logType string, maxSize, maxBackup, maxAge int, compress, debug bool)

	// Level 配置日志级别
	Level() string
}
