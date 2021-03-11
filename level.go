package log

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levels = make(map[string]zap.AtomicLevel)

func getLevel(lvl string) zapcore.Level {
	lvl = strings.ToLower(lvl)

	var level zapcore.Level
	switch lvl {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.ErrorLevel
	}

	return level
}

func SetLevel(system, level string) error {
	if system == "*" {
		for _, s := range levels {
			s.SetLevel(getLevel(level))
		}
		return nil
	}

	if _, exist := levels[system]; !exist {
		return ErrLoggerNotExist
	}

	levels[system].SetLevel(getLevel(level))

	return nil
}
