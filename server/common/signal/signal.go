package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Ignore(f func()) {
	signal.Ignore()
}

func Hup(f func()) {

	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGHUP)

	go func() {
		for {
			<-sighup
			f()
		}
	}()
}

func Int(f func()) {

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)

	go func() {
		<-sigint
		f()
	}()
}

func Quit(f func()) {

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, syscall.SIGQUIT)

	go func() {
		<-sigquit
		f()
	}()

}

func Term(f func()) {

	sigterm := make(chan os.Signal, 1)

	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		<-sigterm
		f()
	}()
}
