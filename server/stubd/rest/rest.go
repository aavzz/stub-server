/*
Package rest implements REST interface of stubd.
*/
package rest

import (
	"github.com/aavzz/stub-server/server/stubd/log"
	"github.com/aavzz/stub-server/server/stubd/rest/api1"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

// InitHttp sets up router.
func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.Handler).Methods("GET")

	if err := http.ListenAndServe(viper.GetString("daemonize"), r); err != nil {
		log.Fatal(err.Error())
	}
}
