package main

/*
 * stdout is unavailable after Setup()
 */

import (
	"net/http"
	. "github.com/aavzz/notifier/setup"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
)

func main() {

	Setup()

	//just some event loop
	//never mind that it does not do anything useful 
	//remove the following two lines and place the server
	//init in Setup()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	RemovePidFile()
}

func handler(w http.ResponseWriter, r *http.Request) {
	SysLog.Info("Hello World!")
}

