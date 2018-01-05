/*
Package api implements rest api version 1 for stubd
*/
package api1

import (
	"encoding/json"
	"github.com/aavzz/daemon/log"
	"github.com/spf13/viper"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	type response struct {
		error int
		errorMsg string
		section string
		key string
		value string
	}

	var resp response
	ret := json.NewEncoder(w)

	section := r.FormValue("section")
	key := r.FormValue("key")

	if section == "" {
		resp.error = 1
		resp.errorMsg = "Section parameter missing or empty"
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}

	if key == "" {
		resp.error = 2
		resp.errorMsg = "Key parameter missing or empty"
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}

	if viper.IsSet(section + "." + key) != true {
		resp.error = 3
		resp.errorMsg = "Key parameter not set"
		resp.section = section
		resp.key = key
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	} else {
		resp.error = 0
		resp.section = section
		resp.key = key
		resp.value = viper.GetString(section + "." + key)
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}
}
