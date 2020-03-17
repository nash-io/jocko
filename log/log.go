package log

import (
	"log"

	"github.com/hashicorp/go-hclog"
)

var (
	Error *log.Logger = hclog.Default().StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: false,
		ForceLevel:  hclog.Error,
	})

	Debug *log.Logger = hclog.Default().StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: false,
		ForceLevel:  hclog.Debug,
	})

	Info *log.Logger = hclog.Default().StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: false,
		ForceLevel:  hclog.Info,
	})

	Trace *log.Logger = hclog.Default().StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: false,
		ForceLevel:  hclog.Trace,
	})

	DebugLevel hclog.Level = hclog.Debug
)

func New(level hclog.Level, prefix string) hclog.Logger {
	l := hclog.New(nil)
	return l
}

func SetPrefix(prefix string) {
	hclog.SetDefault(hclog.Default().ResetNamed(prefix))
}

func SetLevel(level string) {
	lvl := hclog.LevelFromString(level)
	hclog.Default().SetLevel(lvl)
}

func NewStdLogger(l hclog.Logger) *log.Logger {
	stdLogger := l.StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: true,
	})
	return stdLogger
}
