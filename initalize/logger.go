package initalize

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

func InitDatabaseLogger() {
	fmt.Println("执行啦")
	encoder := getEncoder()
	sync := getLogWriter()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	Log = zap.New(core, zap.AddCaller()).Sugar()
	defer Log.Sync()
}

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 负责日志写入的位置
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("Zap.Filename"),
		MaxSize:    viper.GetInt("Zap.MaxSize"),
		MaxBackups: viper.GetInt("Zap.MaxBackups"),
		MaxAge:     viper.GetInt("Zap.MaxAge"),
		Compress:   viper.GetBool("Zap.Compress"),
	}
	syncFile := zapcore.AddSync(lumberJackLogger)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
