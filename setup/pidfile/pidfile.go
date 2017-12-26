package pidfile

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"github.com/tabalt/pidfile"
	. "github.com/aavzz/stub-server/setup/syslog"
	. "github.com/aavzz/stub-server/setup/cmdlnopts"
)

var p *pidfile.PidFile

func WritePid() {
	p = pidfile.NewPidFile(cmdLnOpts.pidfile)
	oldpid, err := p.ReadPidFromFile(p.File)
	if err == nil && oldpid.ProcessExist() {
		SysLog.Err("Another process is already running")
		os.Exit(1)
	}

	//avoid creating pidfile in parent
	daemonState := os.Getenv("_STUB_SERVER_DAEMON_STATE")
	if daemonState == "" {
		if err := p.WritePidToFile(p.File, p.Pid); err != nil {
			SysLog.Err(err.Error())
			os.Exit(1)
		}
	}
}

func RemovePidFile() {
	p.Clear()
}
