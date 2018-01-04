package common

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func spawnChild() (*os.Process, error) {

	daemonState := os.Getenv("_DAEMON_STATE")
	switch daemonState {
	case "":
		syscall.Umask(0022)
		syscall.Setsid()
		os.Setenv("_DAEMON_STATE", "1")
	case "1":
		os.Setenv("_DAEMON_STATE", "")
		return nil, nil
	}

	var attrs os.ProcAttr
	f, err := os.Open("/dev/null")
	if err != nil {
		return nil, err
	}
	attrs.Files = []*os.File{f, f, f}

	exec_path, err := os.Executable()
	if err != nil {
		return nil, err
	}

	p, err := os.StartProcess(exec_path, os.Args, &attrs)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Daemonize spawns a child process and stops parent
func Daemonize() {

	//create child process
	p, err := spawnChild()
	if err != nil {
		log.Fatal(err.Error())
	}

	//only entered in parent
	if p != nil {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)
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
}
