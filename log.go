package bingo

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"github.com/xylong/bingo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

const (
	hourMinutes = 60
	dayMinutes  = hourMinutes * 24
)

var logger *zap.Logger

// logConfig 日志配置
type logConfig struct {
	Level      string
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
	Json       bool
	Duration   int // 时间间隔
}

func newLogConfig() *logConfig {
	config := new(logConfig)
	if err := viper.UnmarshalKey("log", config); err != nil {
		panic(err)
	}

	return config
}

// isDebug 是否为调试模式
func (c *logConfig) isDebug() bool {
	return gin.Mode() == gin.DebugMode
}

// getLevel 获取日志级别
func (c *logConfig) getLevel() zap.AtomicLevel {
	level := zap.NewAtomicLevel()

	// 日志级别（调试模式下日志为调试级别）
	if c.isDebug() == false {
		switch c.Level {
		case "info":
			level.SetLevel(zapcore.InfoLevel)
		case "warn":
			level.SetLevel(zapcore.WarnLevel)
		case "error":
			level.SetLevel(zapcore.ErrorLevel)
		default:
			level.SetLevel(zap.DebugLevel)
		}
	} else {
		level.SetLevel(zap.DebugLevel)
	}

	return level
}

// Zap 初始化日志
func Zap() {
	config := newLogConfig()
	writer := getLogWriter(config)
	encoder := getEncoder(config.Json)

	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writer...), config.getLevel())
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)) // ⚠️级别开始记录堆栈信息
	zap.ReplaceGlobals(logger)                                                // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()、zap.S()调用即可
}

// getEncoder 日志格式配置
func getEncoder(jsonFormat bool) zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}

	// 设置日志编码格式json、console
	if jsonFormat {
		return zapcore.NewJSONEncoder(config)
	}

	return zapcore.NewConsoleEncoder(config)
}

// getLogWriter 日志写入
func getLogWriter(config *logConfig) []zapcore.WriteSyncer {
	var (
		writer   io.Writer
		filename string
	)

	if index := strings.Index(config.FileName, "."); index != -1 {
		var format string

		switch duration := time.Duration(config.Duration / 60); {
		case duration >= dayMinutes:
			format = "%Y%m%d"
		case duration == hourMinutes:
			format = "%Y%m%d%H"
		default:
			format = "%Y%m%d" + utils.ZeroFill(strconv.Itoa(utils.Ceil(time.Now().Hour()*60, config.Duration)), 3, true)
		}

		filename = config.FileName[0:index] + format + config.FileName[index:]
	}

	if config.Duration > 0 {
		writer, _ = rotatelogs.New(
			filename,
			rotatelogs.WithMaxAge(time.Duration(config.MaxSize)*time.Hour),
			rotatelogs.WithRotationTime(time.Minute*time.Duration(config.Duration)),
		)
	} else {
		writer = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
	}

	// 日志分割
	syncer := []zapcore.WriteSyncer{zapcore.AddSync(writer)}
	// 如果是debug模式输出到控制台
	if config.isDebug() {
		syncer = append(syncer, zapcore.AddSync(os.Stdout))
	}

	return syncer
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
