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

	type JResponse struct {
		Error int
		ErrorMsg string
		Section string
		Key string
		Value string
	}

	var resp JResponse
	ret := json.NewEncoder(w)

	section := r.FormValue("section")
	key := r.FormValue("key")

	if section == "" {
		resp.Error = 1
		resp.ErrorMsg = "Section parameter missing or empty"
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}

	if key == "" {
		resp.Error = 2
		resp.ErrorMsg = "Key parameter missing or empty"
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}

	if viper.IsSet(section + "." + key) != true {
		resp.Error = 3
		resp.ErrorMsg = "Key parameter not set"
		resp.Section = section
		resp.Key = key
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	} else {
		resp.Error = 0
		resp.Section = section
		resp.Key = key
		resp.Value = viper.GetString(section + "." + key)
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}
}
