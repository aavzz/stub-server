/*
stubd does nothing but starts, reads command line options,
initializes logging, sets up signal handling, writes PID file
and sits in the background.

This daemon demonstrates the concept of daemon process creation
in Go and can be used as a starting point for other projects
that do something usefull
*/
package main

import (
	. "github.com/aavzz/stub-server/setup"
	. "github.com/aavzz/stub-server/setup/pidfile"
)

func main() {
	Setup()
	RemovePidFile()
}
