package parrot

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New("test", OptSetLevel("ERROR"), OptSetLogFile("log", 1, 0, 1))

	logger.Info(1)
	logger.Warn(1)
	logger.Error(1)
	logger.Debug(1)

	_ = SetLevel("test", "DEBUG") // change log level

	logger.Info(2)
	logger.Warn(2)
	logger.Error(2)
	logger.Debug(2)
}

func TestLogToFile(t *testing.T) {
	logger := New("file-logger", OptSetLevel("DEBUG"), OptSetLogFile("log", 1, 0, 1))

	logger.Info(1)
	logger.Warn(1)
	logger.Error(1)
	logger.Debug(1)
}
