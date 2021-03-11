package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/xerrors"
)

type config struct {
	w   io.Writer
	lvl zapcore.Level
}

var ErrLoggerNotExist = xerrors.New("logger does not exist")

func defaultConf() config {
	return config{
		w:   os.Stdout,
		lvl: zapcore.ErrorLevel,
	}
}

func New(system string, opts ...Option) *zap.SugaredLogger {
	conf := defaultConf()

	for _, o := range opts {
		o(&conf)
	}

	if len(system) == 0 {
		system = "undefined"
	}

	if _, exist := levels[system]; !exist {
		level := zap.NewAtomicLevel()
		level.SetLevel(conf.lvl)
		levels[system] = level
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(newEncoderConfig(false)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(conf.w)),
		levels[system],
	)

	logger := zap.New(core).WithOptions(zap.IncreaseLevel(levels[system]), zap.AddCaller()).Named(system).Sugar()

	return logger
}

func newEncoderConfig(caller bool) zapcore.EncoderConfig {
	enc := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	if caller {
		enc.CallerKey = "caller"
	}

	return enc
}
