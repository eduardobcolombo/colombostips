package log

type Config struct {
	Level string
}

func New(cfg Config) string {
	return cfg.Level
}
