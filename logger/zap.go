package logger

import (
	"github/go-sven/sven-layout/app/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	sugar  *zap.SugaredLogger
	Logger *zap.Logger
	TraceKey = "traceId"
)

func NewZapLogger(config *conf.LoggerConfig) {
	writeSyncer := getLogWriter(config)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// zap打印时将整个调用的stack链路会存放到内存中，默认打印调用处的caller信息。所以需要再初始化zap时额外增加AddCallerSkip跳过指定层级的caller
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer Logger.Sync()
	sugar = Logger.Sugar()
}

// getEncoder zapcore.Encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(config *conf.LoggerConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Filename, // 日志文件位置
		MaxSize:    config.MaxSize,                                 // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxBackups:	config.MaxBackups,                              // 保留旧文件的最大个数
		MaxAge:     config.MaxAge,                                  // 保留旧文件的最大天数
		Compress:   false,                                // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}