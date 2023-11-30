package logger

// Config has all needed values to create the Logger.
type Config struct {
	Level string
}

// Logger has all is needed to use Logger pkg.
type Logger struct {
	Level string // hack for this demo
}

// New provides a Logger object as a Factory pattern.
func New(cfg Config) *Logger {
	return &Logger{
		Level: cfg.Level,
	}
}
