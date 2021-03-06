package logger

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
)

const (
	EncoderConsole = "console"
	EncoderJson    = "json"
)

const (
	DriverZap string = "zap"
)

var (
	ErrEmptyLoggerConfig   = errors.New("empty logger config")
	ErrNoValidLoggerConfig = errors.New("no valid logger config")
)

var log Logger

type Fields map[string]interface{}

var timeFormat = "2006-01-02T15:04:05.000Z0700"

type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Warnf(format string, args ...interface{})
	Warn(args ...interface{})

	Errorf(format string, args ...interface{})
	Error(args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})

	Panicf(format string, args ...interface{})
	Panic(args ...interface{})

	WithFields(keyValues Fields) Logger
}

func NewLogger(v *viper.Viper, driver string) error {
	l, err := GetLogger(v, driver)
	if err != nil {
		return err
	}

	log = l
	return nil
}

func NewDefaultLogger() error {
	l, err := newDefaultZapLogger()
	if err != nil {
		return err
	}

	log = l
	return nil
}

func GetLogger(v *viper.Viper, driver string) (Logger, error) {
	switch driver {
	case DriverZap:
		return newZapLogger(v)
	default:
		return newDefaultZapLogger()
	}
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}

func keys(v *viper.Viper) []string {
	a := v.AllSettings()
	keys := make([]string, len(a))
	i := 0
	for k := range a {
		keys[i] = k
		i++
	}
	return keys
}
