package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string
}

func New(cfg Config) *zap.Logger {

	zapEncoder := zap.NewProductionEncoderConfig()
	zapEncoder.TimeKey = "timestamp"
	zapEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	cfgLog := zap.Config{
		Level:         logLevel(cfg.Level),
		Encoding:      "json",
		EncoderConfig: zapEncoder,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger, err := cfgLog.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func logLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	default:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	}
}
