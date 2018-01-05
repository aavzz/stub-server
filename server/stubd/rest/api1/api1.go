/*
Package api implements rest api version 1 for stubd
*/
package api1

import (
	"github.com/aavzz/daemon/log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Hello World!")
}
