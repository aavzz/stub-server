package pid

import (
	"github.com/aavzz/stub-server/server/common/log"
	"github.com/tabalt/pidfile"
)

var p *pidfile.PidFile

// WritePid checks if the process has already been started
// and writes PID file
func Write(pid string) {
	p = pidfile.NewPidFile(pid)
	oldpid, err := p.ReadPidFromFile(p.File)
	if err == nil && oldpid.ProcessExist() {
		log.Fatal("Another process is already running")
	}

	if err := p.WritePidToFile(p.File, p.Pid); err != nil {
		log.Fatal(err.Error())
	}
}

func Remove() {
	if p != nil {
		p.Clear()
	}
}
