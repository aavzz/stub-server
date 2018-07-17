/*
Package api implements rest api version 2 for stubd
*/
package api2

import (
	"encoding/json"
	"github.com/aavzz/daemon/log"
	"github.com/spf13/viper"
	"net/http"
	"io/ioutil"
	"crypto/sha256"
	"regexp"
	"strconv"
)

// Handler processes http requests and writes responses
func Handler(w http.ResponseWriter, r *http.Request) {


	var resp JResponse
	ret := json.NewEncoder(w)

	body, err := ioutil.ReadAll(r.Body)
        if err != nil {
		log.Error(err.Error())
                return 
        }
        var v JRequest
        if err := json.Unmarshal(body, &v); err != nil {
		resp.Error = 1
		resp.ErrorMsg = err.Error()
		resp.Signature = 
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}

	//check signature

	signature := JRequest.Signature
	JRequest.Signature = viper.GetString("stubd.secret")
	request, err := json.Marshall(JRequest)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if string(sha256.Sum256(request)) != signature {
		resp.Error = 1
		resp.ErrorMsg = "Signature check failed"
		resp.Signature = 
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
		
	}

	//signature ok, check user input

	if m, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-A0-9_-]*$`, JRequest.Section); !m {
		resp.Error = 1
		resp.ErrorMsg = "Wrong section name"
		resp.Signature = 
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
                return
        }

	if m, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-A0-9_-]*$`, JRequest.Key); !m {
		resp.Error = 1
		resp.ErrorMsg = "Wrong key name"
		resp.Signature = 
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
                return
        }

	if viper.IsSet(JRequest.Section+"."+JRequest.Key) != true {
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
		resp.Section = JRequest.Section
		resp.Key = JRequest.Key
		if JRequest.Section == "stubd" &&  JRequest.Key == "secret" {
			resp.Value = "** CENSORED **"
		} else {
			resp.Value = viper.GetString(JRequest.Section + "." + JRequest.Key)
		}
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
		return
	}
}
