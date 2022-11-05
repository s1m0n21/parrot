package parrot

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/xerrors"
)

type config struct {
	w      io.Writer
	lvl    zapcore.Level
	caller bool
	color  bool
}

var ErrLoggerNotExist = xerrors.New("logger does not exist")

func defaultConf() *config {
	return &config{
		w:      os.Stdout,
		lvl:    zapcore.ErrorLevel,
		caller: true,
		color:  true,
	}
}

func New(system string, opts ...Option) *zap.SugaredLogger {
	conf := defaultConf()

	for _, o := range opts {
		o(conf)
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
		zapcore.NewConsoleEncoder(newEncoderConfig(conf)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(conf.w)),
		levels[system],
	)

	logger := zap.New(core).WithOptions(zap.IncreaseLevel(levels[system]), zap.AddCaller()).Named(system).Sugar()

	return logger
}

func newEncoderConfig(conf *config) zapcore.EncoderConfig {
	enc := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	if conf.caller {
		enc.CallerKey = "caller"
	}

	if conf.color {
		enc.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return enc
}
