package logs

import (
	"common/config"
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var logger *log.Logger

func InitLog(appName string) {
	logger = log.New(os.Stderr)

	if config.Conf.Log.Level == "DEBUG" {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetLevel(log.InfoLevel)
	}
	logger.SetPrefix(appName)
	logger.SetReportTimestamp(true)
	logger.SetTimeFormat(time.DateTime)
	//logger.SetLevel(log.InfoLevel)
}

func Fatal(format string, args ...any) {
    if (len(args) == 0) {
        logger.Fatal(format)
	}else{
		logger.Fatalf(format, args...)
	}
}

func Info(format string, args ...any) {
    if (len(args) == 0) {
        logger.Info(format)
	}else{
		logger.Infof(format, args...)
	}
}

func Warn(format string, args ...any) {
    if (len(args) == 0) {
        logger.Warn(format)
	}else{
		logger.Warnf(format, args...)
	}
}

func Debug(format string, args ...any) {
    if (len(args) == 0) {
        logger.Debug(format)
	}else{
		logger.Debugf(format, args...)
	}
}

func Error(format string, args ...any) {
    if (len(args) == 0) {
        logger.Error(format)
	}else{
		logger.Errorf(format, args...)
	}
}