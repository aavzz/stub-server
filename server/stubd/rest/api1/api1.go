/*
Package runtime/stubd implements everything than makes
stubd usefull.
*/
package api1

import (
	"github.com/aavzz/stub-server/server/common/log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Hello World!")
}
