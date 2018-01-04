package cmd

import (
	"github.com/aavzz/stub-server/server/common"
	"github.com/aavzz/stub-server/server/stubd/log"
	"github.com/aavzz/stub-server/server/stubd/pid"
	"github.com/aavzz/stub-server/server/stubd/signal"
	"github.com/aavzz/stub-server/server/stubd/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stubd = &cobra.Command{
	Use:   "stubd",
	Short: "stubd is a minimal example of a rest web-app",
	Long:  `A minimal web-application to use as a base for larger projects`,
	Run:   stubdCommand,
}

func stubdCommand(cmd *cobra.Command, args []string) {

	if viper.GetBool("daemonize") == true {
		log.InitSyslog()
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
		signal.Handling()
	}
	rest.InitHttp()
}

func Execute() {
	stubd.Flags().StringP("config", "c", "/etc/stubd.conf", "configuration file (default: /etc/stubd.conf)")
	stubd.Flags().StringP("pidfile", "p", "/var/run/stubd.pid", "PID file (default: /var/run/stubd.pid)")
	stubd.Flags().StringP("address", "a", "127.0.0.1:8082", "address and port to bind to (default: 127.0.0.1:8082)")
	stubd.Flags().BoolP("daemonize", "d", false, "Run as a daemon (default: no)")
	viper.BindPFlag("config", stubd.Flags().Lookup("config"))
	viper.BindPFlag("pidfile", stubd.Flags().Lookup("pidfile"))
	viper.BindPFlag("address", stubd.Flags().Lookup("address"))
	viper.BindPFlag("daemonize", stubd.Flags().Lookup("daemonize"))

	if err := stubd.Execute(); err != nil {
		log.Fatal(err.Error())
	}

}
