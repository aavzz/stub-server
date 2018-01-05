/*
Package rest implements REST interface of stubd.
*/
package rest

import (
	"context"
	"github.com/aavzz/stub-server/server/common/log"
	"github.com/aavzz/stub-server/server/common/pid"
	"github.com/aavzz/stub-server/server/common/signal"
	"github.com/aavzz/stub-server/server/stubd/rest/api1"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// InitHttp sets up router.
func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.Handler).Methods("GET")

	s := &http.Server{
		Addr:     viper.GetString("address"),
		Handler:  r,
		ErrorLog: log.Logger("stubd"),
	}

	if viper.GetBool("daemonize") == true {
		signal.Quit(func() {
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			log.Info("SIGQUIT received, exitting gracefully")
			s.Shutdown(ctx)
			pid.Remove()
		})
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
