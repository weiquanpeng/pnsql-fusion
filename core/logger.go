package core

import (
	"github.com/xiaohongshu/PnSql/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if global.P_cfg.Zap.Level == "" {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)

}

func getWriteSyncer() zapcore.WriteSyncer {
	s := string(filepath.Separator)
	st, _ := os.Getwd()
	stLogPathFile := filepath.Join(st, "logs", s, time.Now().Format(time.DateOnly)+".log")

	luBerJackSyncer := &lumberjack.Logger{
		Filename:   stLogPathFile,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	return zapcore.AddSync(luBerJackSyncer)
}
