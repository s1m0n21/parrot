package parrot

import (
	"os"

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

func OptLevelFromEnv() Option {
	return func(c *config) {
		if lvl, set := os.LookupEnv("PARROT_LEVEL"); set {
			c.lvl = getLevel(lvl)
		}
	}
}

func OptWithoutCall() Option {
	return func(c *config) {
		c.caller = false
	}
}

func OptWithoutColor() Option {
	return func(c *config) {
		c.color = false
	}
}
