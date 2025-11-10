package logger

import (
	"go-demo3/internal/config"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 自定义日志级别显示（全大写）
func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	levelStr := strings.ToUpper(level.String())
	enc.AppendString(levelStr)
}

// 自定义时间编码器，格式为"2006-01-02 15:04:05.000"
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 自定义caller编码器，以[]包裹
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	if caller.Defined {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	} else {
		enc.AppendString("[]")
	}
}

func NewLogger(cfg *config.LogConfig) *zap.Logger {
	// lumberjack进行切割
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.MaxSize, // MB
		MaxAge:     cfg.MaxAge,  // days
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	})

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customLevelEncoder,  // 自定义级别显示
		EncodeTime:     customTimeEncoder,   // 自定义时间格式
		EncodeCaller:   customCallerEncoder, // 自定义调用者格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		level = zapcore.InfoLevel // 默认info
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		level,
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}
