package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"wat.ink/layout/fiber/pkg/config"
)

var log *zap.Logger

func init() {
	// 创建基础配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建日志轮转配置
	hook := &lumberjack.Logger{
		Filename:   config.Conf.Log.Filename,
		MaxSize:    config.Conf.Log.MaxSize,
		MaxBackups: config.Conf.Log.MaxBackups,
		MaxAge:     config.Conf.Log.MaxAge,
		Compress:   config.Conf.Log.Compress,
	}

	// 设置日志级别
	var level zapcore.Level
	switch config.Conf.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		level,
	)

	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Info 记录 info 级别的日志
func Info(msg string, keysAndValues ...interface{}) {
	log.Sugar().Infow(msg, keysAndValues...)
}

// Error 记录 error 级别的日志
func Error(msg string, keysAndValues ...interface{}) {
	log.Sugar().Errorw(msg, keysAndValues...)
}

// Debug 记录 debug 级别的日志
func Debug(msg string, keysAndValues ...interface{}) {
	log.Sugar().Debugw(msg, keysAndValues...)
}

// Warn 记录 warn 级别的日志
func Warn(msg string, keysAndValues ...interface{}) {
	log.Sugar().Warnw(msg, keysAndValues...)
}

// Fatal 记录 fatal 级别的日志
func Fatal(msg string, keysAndValues ...interface{}) {
	log.Sugar().Fatalw(msg, keysAndValues...)
} 