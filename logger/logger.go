package logger

import (
	"sync"

	"go.uber.org/zap"

	"url-shortener/config"
)

// In contexts where performance is nice, but not critical, use the SugaredLogger
type Logger struct {
	*zap.SugaredLogger
}

var (
	logger Logger
	once   sync.Once
)

func Get() *Logger {
	once.Do(
		func() {
			cfg := config.Get()

			var zapConfig zap.Config
			if cfg.LogLevel == "debug" {
				zapConfig = zap.NewDevelopmentConfig()
			} else {
				zapConfig = zap.NewProductionConfig()
			}

			switch cfg.LogLevel {
			case "debug":
				zapConfig.Level.SetLevel(zap.DebugLevel)
			case "info":
				zapConfig.Level.SetLevel(zap.InfoLevel)
			case "warn", "warning":
				zapConfig.Level.SetLevel(zap.WarnLevel)
			case "err", "error":
				zapConfig.Level.SetLevel(zap.ErrorLevel)
			case "fatal":
				zapConfig.Level.SetLevel(zap.FatalLevel)
			case "panic":
				zapConfig.Level.SetLevel(zap.PanicLevel)
			default:
				zapConfig.Level.SetLevel(zap.InfoLevel) // log info and above by default
			}

			zeroLogger, _ := zapConfig.Build()
			defer zeroLogger.Sync()
			logger = Logger{zeroLogger.Sugar()}
		},
	)
	return &logger
}
