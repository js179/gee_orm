package logf

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[Erro]\033[0m", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[Info]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Panic  = errorLog.Panic
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	infoLevel = iota
	ErrorLevel
	All
)

func SetLevel(level int) {
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if level >= All {
		return
	} else if level > ErrorLevel {
		errorLog.SetOutput(io.Discard)
	} else if level > infoLevel {
		infoLog.SetOutput(io.Discard)
	}
}
