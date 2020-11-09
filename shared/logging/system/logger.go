package system

import (
	"log"
	"os"
)

type (
	SysLogger struct {
		Logger    *log.Logger
		isEnabled bool
	}
)

func NewSysLogger(enabled bool) *SysLogger {
	return &SysLogger{
		Logger:    log.New(os.Stderr, "[SYSTEM]", log.Ldate|log.Ltime|log.Lshortfile),
		isEnabled: enabled,
	}
}

func (logger *SysLogger) Log(arg ...interface{}) {
	if logger.isEnabled {
		logger.Logger.Print(arg...)
	}
}
