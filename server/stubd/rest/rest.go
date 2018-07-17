/*
Package rest implements REST interface of stubd
*/
package rest

import (
	"github.com/aavzz/daemon/log"
	"github.com/aavzz/daemon/pid"
	"github.com/aavzz/daemon/signal"
	"github.com/aavzz/stub-server/server/stubd/rest/api1"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var Server *http.Server

// InitHttp sets up router and starts server
func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.Handler).Methods("POST")

	Server = &http.Server{
		Addr:     viper.GetString("address"),
		Handler:  r,
		ErrorLog: log.Logger("stubd"),
	}

	go func() {
		if err := S.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()
}
