package application

import (
	"log"
	"os"
)

type (
	ApplicationLogger interface {
		LogWithTimeDiff(arg ...interface{})
	}

	AppLogger struct {
		Logger  *log.Logger
		Service *interface{}
	}
)

func NewAppLogger(service *interface{}) *AppLogger {
	return &AppLogger{
		Logger:  log.New(os.Stderr, "[SYSTEM]", log.Ldate|log.Ltime|log.Lshortfile),
		Service: service,
	}
}
