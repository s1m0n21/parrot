package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option func(*config)

func OptSetLogFile(filename string, size, backups, age int) Option {
	return func(c *config) {
		c.w = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    size, // megabytes
			MaxBackups: backups,
			MaxAge:     age, // days
			Compress:   true,
		}
	}
}

func OptSetLevel(lvl string) Option {
	return func(c *config) {
		c.lvl = getLevel(lvl)
	}
}
