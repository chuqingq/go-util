package util

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger
var defaultLoggerOnce sync.Once

func Logger() *logrus.Logger {
	defaultLoggerOnce.Do(func() {
		logger = &logrus.Logger{
			Out: &lumberjack.Logger{
				Filename:   "/dev/shm/test.log", // in memory
				MaxSize:    1,                   // megabytes
				MaxBackups: 1,                   // reserve 1 backup
				// MaxAge:     28, //days
			},
			// Formatter: &logrus.TextFormatter{},
			Formatter: &myFormatter{},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		}
		logger.SetReportCaller(true)
		// defaultLogger.SetLevel(logrus.WarnLevel)
	})
	return logger
}

type myFormatter struct {
}

func splitAndGetLast(str string, sep string) string {
	slice := strings.Split(str, sep)
	return slice[len(slice)-1]
}

func (f *myFormatter) Format(e *logrus.Entry) ([]byte, error) {
	// time level caller message
	return []byte(fmt.Sprintf("%s %5.5s [%s:%v %s] %s\n",
		e.Time.Format("01/02 15:04:05.000000"),
		e.Level.String(),
		splitAndGetLast(e.Caller.File, "/"),
		e.Caller.Line, splitAndGetLast(e.Caller.Function, "."),
		e.Message)), nil
}

// TODO
// 1.output, appendoutput, removealloutputs
// 2.log file name
// 3.log level
