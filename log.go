package bingo

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var logger *zap.Logger

type LogConfig struct {
	Level      string
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

// getLevel 获取日志级别
func (c *LogConfig) getLevel() zap.AtomicLevel {
	level, mode := zap.NewAtomicLevel(), gin.DebugMode

	switch mode {
	case "debug":
		level.SetLevel(zapcore.DebugLevel)
	case "info":
		level.SetLevel(zapcore.InfoLevel)
	case "warn":
		level.SetLevel(zapcore.WarnLevel)
	case "error":
		level.SetLevel(zapcore.ErrorLevel)
	default:
		level.SetLevel(zap.DebugLevel)
	}

	return level
}

// isDebugLevel 判断是否为debug级别
func (c *LogConfig) isDebugLevel() bool {
	return c.Level == gin.DebugMode
}

func InitLog(config *LogConfig) error {
	writer := getLogWriter(config)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writer...), config.getLevel())
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)) // ⚠️级别开始记录堆栈信息

	zap.ReplaceGlobals(logger)
	return nil
}

// getEncoder 日志格式配置
func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006/01/02 15:04:05.000"))
	}

	// 设置日志编码格式json、console
	if viper.GetString("log.mode") == "json" {
		return zapcore.NewJSONEncoder(config)
	}

	return zapcore.NewConsoleEncoder(config)
}

// getLogWriter 日志写入
func getLogWriter(config *LogConfig) []zapcore.WriteSyncer {
	// 日志分割
	hook := &lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	writer := []zapcore.WriteSyncer{zapcore.AddSync(hook)}
	// 如果是debug模式输出到控制台
	if config.isDebugLevel() {
		writer = append(writer, zapcore.AddSync(os.Stdout))
	}

	return writer
}

// GinRecovery 代替gin的默认recovery
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool

				// 检查连接是否断开
				if v, ok := err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(v.Error()), "broken pipe") || strings.Contains(strings.ToLower(v.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}

				httpRequest, _ := httputil.DumpRequest(context.Request, false)
				if brokenPipe {
					logger.Error(context.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)

					context.Error(err.(error))
					context.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}

				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		context.Next()
	}
}
