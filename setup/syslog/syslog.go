/*
Package syslog sets up logging to local syslog
*/
package syslog

/*
 * this code runs both in parent and child
 * parent's output goes to stdout,
 * child's to /dev/null
 */

import (
	"fmt"
	"log/syslog"
	"os"
)

// SysLog methods do the logging
var SysLog *syslog.Writer

// InitLogging tries to get in touch with local syslog.
// It writes error message to stdout and stops the
// process in case of failure.
func InitLogging() {
	var err error
	SysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "stubd")
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}
