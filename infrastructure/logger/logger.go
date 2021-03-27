package logger

import (
	"log"
	"os"
)

type (
	logger struct {
		err *log.Logger
		out *log.Logger
	}
)

func NewLogger() *logger {
	return &logger{
		err: log.New(os.Stderr, "", 0),
		out: log.New(os.Stdout, "", 0),
	}
}

func (l *logger) Info(format string, v ...interface{}) {
	l.out.Printf(format, v...)
}

func (l *logger) Error(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
