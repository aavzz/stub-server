/*
Package setup initializes the process to run in the
background.
*/
package setup

import (
	"fmt"
	. "github.com/aavzz/stub-server/setup/cfgfile"
	. "github.com/aavzz/stub-server/setup/cmdlnopts"
	. "github.com/aavzz/stub-server/setup/daemon"
	. "github.com/aavzz/stub-server/setup/pidfile"
	. "github.com/aavzz/stub-server/setup/signal"
	. "github.com/aavzz/stub-server/setup/syslog"
	. "github.com/aavzz/stub-server/serverlogic/serverlogic"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Setup spawns child process, checks that everything is ok,
// does the necessary initialization and stops the parent process.
func Setup() {

	//create child process
	p, err := Daemonize()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//only entered in parent
	if p != nil {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)

		//checks command line args and tries to init logging
		//writes errors to stdout
		ParseCmdLine()
		InitLogging()

		<-sigterm
		os.Exit(0)
	}

	//parent never gets here

	//give the parent time to install signal handler
	//so we don't kill it prematurely and get ugly message on stdout
	time.Sleep(100 * time.Millisecond)

	//say good bye to parent
	ppid := os.Getppid()
	//we don't want to kill init
	if ppid > 1 {
		syscall.Kill(ppid, syscall.SIGTERM)
	}

	//child's output goes to /dev/null
	//we call this in parent just to check for correctness
	//and see error indications on stdout (logging is not initialized yet)
	//real configuration happens here
	ParseCmdLine()
	InitLogging()

	//final touch
	WritePid()
	ReadConfig()
	SignalHandling()
	ServerLogicInit()
}
