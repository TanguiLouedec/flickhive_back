package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zap.DebugLevel)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func GetLogger() *zap.Logger {
	return Logger
}
