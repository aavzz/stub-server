package syslog

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"fmt"
	"log/syslog"
)

var SysLog *syslog.Writer

func initLogging() {
	var err error
	SysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "stub-server")
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}
