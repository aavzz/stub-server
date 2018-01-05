package log

import (
	"log"
	"log/syslog"
	"os"
)

var sysLog *syslog.Writer

func Logger(tag string) *log.Logger {
	l := log.New(os.Stdout, tag, syslog.LOG_INFO|syslog.LOG_DAEMON)
	if sysLog != nil {
		l.SetOutput(sysLog)
	}
	return l
}

func InitSyslog(tag string) {
	var err error
	sysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, tag)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Info(s string) {
	if sysLog != nil {
		sysLog.Info(s)
	} else {
		log.Print(s)
	}
}

func Error(s string) {
	if sysLog != nil {
		sysLog.Err(s)
	} else {
		log.Print(s)
	}
}

func Fatal(s string) {
	if sysLog != nil {
		sysLog.Err(s)
		os.Exit(1)
	} else {
		log.Print(s) //log.Fatal produces ugly syslog message
		os.Exit(1)
	}
}
