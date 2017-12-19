package setup

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"github.com/pborman/getopt/v2"
)

type commandLineOptions struct {
	cfgfile string
	pidfile string
}

var cmdLnOpts commandLineOptions

func parseCmdLine() {
	c := getopt.StringLong("cfg", 'c', "/etc/notifyd.conf", "configuration file")
	p := getopt.StringLong("pid", 'p', "/var/run/notifyd.pid", "PID file")
	getopt.Parse()
	cmdLnOpts.cfgfile = *c
	cmdLnOpts.pidfile = *p
}
