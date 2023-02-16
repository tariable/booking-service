package logger

import (
	"bookingService/config"
	"fmt"
	"log"
)

type Logger interface {
	Info(msg string)
	Infof(format string, v ...any)
	Errorf(format string, v ...any)
}

type DefaultLogger struct {
	l *log.Logger
}

func (log *DefaultLogger) Info(msg string) {
	log.l.Printf("[Info]: %s\n", msg)
}

func (log *DefaultLogger) Infof(format string, v ...any) {
	msg := fmt.Sprintf(format, v)
	log.l.Printf("[Info]: %s\n", msg)
}

func (log *DefaultLogger) Errorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v)
	log.l.Printf("[Error]: %s\n", msg)
}

func New(config config.Logger) Logger {
	return &DefaultLogger{l: log.Default()}
}
