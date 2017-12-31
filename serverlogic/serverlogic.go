/*
Package serverlogick implements everything than makes
the server usefull. Replace this package with whatever
you think appropriate.
*/
package serverlogic

import (
	. "github.com/aavzz/stub-server/setup/syslog"
	"net/http"
)

// ServerLogickInit initializes whatever you want yout server to do.
// Replace it with the right routine for your project.
func ServerLogicInit() {

	//just some event loop
	//never mind that it does not do anything useful
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	SysLog.Info("Hello World!")
}
