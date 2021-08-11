package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func CreateLogger() *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("could not initialize zap logger: %v", err)
	}
	return logger.Sugar()
}
