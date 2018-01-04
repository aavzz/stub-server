package signal

import (
	"github.com/aavzz/stub-server/server/stubd/log"
	"github.com/aavzz/stub-server/server/stubd/pid"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Handling() {

	sighup := make(chan os.Signal, 1)
	sigint := make(chan os.Signal, 1)
	sigquit := make(chan os.Signal, 1)
	sigterm := make(chan os.Signal, 1)

	signal.Notify(sighup, syscall.SIGHUP)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigquit, syscall.SIGQUIT)
	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		for {
			<-sighup
			log.Info("SIGHUP received, re-reading configuration file")
			if err := viper.ReadInConfig(); err != nil {
				pid.Remove()
				log.Fatal(err.Error())
			}
		}
	}()

	go func() {
		<-sigint
		log.Info("SIGINT received, exitting")
		pid.Remove()
		os.Exit(0)
	}()

	go func() {
		<-sigquit
		log.Info("SIGQUIT received, exitting")
		pid.Remove()
		os.Exit(0)
	}()

	go func() {
		<-sigterm
		log.Info("SIGTERM received, exitting")
		pid.Remove()
		os.Exit(0)
	}()
}
