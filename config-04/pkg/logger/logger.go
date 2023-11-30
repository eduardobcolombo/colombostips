package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config has all needed values to create the Logger.
type Config struct {
	Level string
}

// Logger has all is needed to use Logger pkg.
type Logger struct {
	Log *zap.Logger
}

// New provides a Logger object as a Factory pattern.
func New(cfg Config) *Logger {

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

	return &Logger{
		Log: logger,
	}
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
