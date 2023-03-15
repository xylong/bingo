package log

import (
	"fmt"
	"github.com/xylong/bingo/iface"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

// LoggerConfig 日志配置
type LoggerConfig struct {
	Filename  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
	LogType   string
	Level     string
}

// InitLogger 日志初始化
func InitLogger(configure iface.Configure) {

	// 获取日志存储格式
	enc := getEncoder(configure.Encoder())
	// 获取日志写入介质
	ws := getWriter(configure.Writer())
	// 设置日志等级
	logLevel := getLevel(configure.Level())

	// 初始化 core
	core := zapcore.NewCore(enc, zapcore.NewMultiWriteSyncer(ws...), logLevel)

	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用自定义的 Logger
	zap.ReplaceGlobals(Logger)
}

// getEncoder 日志存储格式
func getEncoder(config zapcore.EncoderConfig, debug bool) zapcore.Encoder {
	if debug {
		// 终端输出的关键词高亮
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 控制台输出（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(config)
	}

	return zapcore.NewJSONEncoder(config)
}

// 日志记录介质。Bingo 中使用了两种介质，os.Stdout 和文件
func getWriter(filename, logType string, maxSize, maxBackup, maxAge int, compress, debug bool) []zapcore.WriteSyncer {
	if logType == "daily" {
		name := time.Now().Format("2006-01-02") + ".log"
		filename = strings.ReplaceAll(filename, "logs.log", name)
	}

	// 滚动日志
	logger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
	}

	// 记录到文件
	writeSyncer := []zapcore.WriteSyncer{zapcore.AddSync(logger)}
	// 调试打印到终端
	if debug {
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	return writeSyncer
}

// getLevel 获取日志级别
func getLevel(level string) zap.AtomicLevel {
	logLevel := zap.NewAtomicLevel()
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误")
	}

	return logLevel
}
