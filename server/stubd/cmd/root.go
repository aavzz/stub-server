/*
Package cmd implements stubd commands and flags
*/
package cmd

import (
	"github.com/aavzz/daemon"
	"github.com/aavzz/daemon/log"
	"github.com/aavzz/daemon/pid"
	"github.com/aavzz/daemon/signal"
	"github.com/aavzz/stub-server/server/stubd/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var stubd = &cobra.Command{
	Use:   "stubd",
	Short: "stubd is a minimal example of a rest web-app",
	Long:  `A minimal web-application to use as a base for larger projects`,
	Run:   stubdCommand,
}

func stubdCommand(cmd *cobra.Command, args []string) {

	if viper.GetBool("daemonize") == true {
		log.InitSyslog("stubd")
		common.Daemonize()
	}

	//After daemonize() this part runs in child only

	viper.SetConfigType("toml")
	viper.SetConfigFile(viper.GetString("config"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if viper.GetBool("daemonize") == true {
		pid.Write(viper.GetString("pidfile"))
		signal.Ignore()
		signal.Hup(func() {
			log.Info("SIGHUP received, re-reading configuration file")
			if err := viper.ReadInConfig(); err != nil {
				pid.Remove()
				log.Fatal(err.Error())
			}
		})
		signal.Term(func() {
			log.Info("SIGTERM received, exitting")
			pid.Remove()
			os.Exit(0)
		})
	}
	rest.InitHttp()
}

// Execute starts stubd execution
func Execute() {
	stubd.Flags().StringP("config", "c", "/etc/stubd.conf", "configuration file (default: /etc/stubd.conf)")
	stubd.Flags().StringP("pidfile", "p", "/var/run/stubd.pid", "PID file (default: /var/run/stubd.pid)")
	stubd.Flags().StringP("address", "a", "127.0.0.1:8082", "address and port to bind to (default: 127.0.0.1:8082)")
	stubd.Flags().BoolP("daemonize", "d", false, "run as a daemon (default: no)")
	viper.BindPFlag("config", stubd.Flags().Lookup("config"))
	viper.BindPFlag("pidfile", stubd.Flags().Lookup("pidfile"))
	viper.BindPFlag("address", stubd.Flags().Lookup("address"))
	viper.BindPFlag("daemonize", stubd.Flags().Lookup("daemonize"))

	if err := stubd.Execute(); err != nil {
		log.Fatal(err.Error())
	}

}
