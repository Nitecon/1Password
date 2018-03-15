package log

import (
	"flag"
	"sync"

	"github.com/Nitecon/1Password/exit"
	"go.uber.org/zap"
)

var (
	logger     *Logger
	loggerLock = new(sync.RWMutex)
	level      = 5
)

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
}

func init() {
	flag.IntVar(&level, "Log level", 5, "Set the log level\n 0: Debug\n5: Info: 7: Warning\n 9: Error")
	flag.Parse()
	log, err := zap.NewProduction()
	if err != nil {
		exit.Fatalm("Could not initialize logger", err)
	}
	loggerLock.Lock()
	defer loggerLock.Unlock()
	*logger = log
}
