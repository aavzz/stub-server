package cmdlnopts

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

func ParseCmdLine() {
	c := getopt.StringLong("cfg", 'c', "/etc/stub-server.conf", "configuration file")
	p := getopt.StringLong("pid", 'p', "/var/run/stub-server.pid", "PID file")
	getopt.Parse()
	cmdLnOpts.cfgfile = *c
	cmdLnOpts.pidfile = *p
}

func ConfigPidFile() string {
	return cmdLnOpts.pidfile
}

func ConfigCfgFile() string {
	return cmdLnOpts.cfgfile
}
