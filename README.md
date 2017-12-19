This is a server which does nothing but starts and sits there.

Well, it does do something:

- detaches itself from terminal
- processes command line paramaters
- reads a config file
- logs to syslog
- reacts to some signals
- checks if it has been already invoked and refuses to start
- writes a pidfile

The problem with Golang daemons is that the traditional way of daemonizing a process through fork() is not available in go.
So another approach is used. A process starts, spawns a terminal-detached and setsid()ed child, then gets killed by the child.
Synchronization between the two is carried out through an environment variable.

There is a number of caveats here:

- All the process setup (see above) happens in both parent and child
- Syslog logging is not available to parent
- Terminal i/o is not available to child
- A race condition when parent gets killed before it has enough time to do checks

All of the above is dealt with in this example stub-server

This software came into being after I studied a number of other daemonizing projects for the Go language.
I did not invent anything new here, just wrapped up the ideas and sample code snippets.

This software can be used as the basis of real golang servers. Drop me a line if you find it useful.
