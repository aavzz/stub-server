/*
Package cmd implements stubc commands and flags
*/
package cmd

import (
	"crypto/tls"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var stubc = &cobra.Command{
	Use:   "stubc",
	Short: "stubc is a minimal example of a rest web-app client",
	Long:  `A minimal web-application client to use as a base for larger projects`,
	Run:   stubcCommand,
}

func stubcCommand(cmd *cobra.Command, args []string) {

	type JResp struct {
		Error    int
		ErrorMsg string
		Section  string
		Key      string
		Value    string
	}

	parameters := url.Values{
		"section": {viper.GetString("section")},
		"key":     {viper.GetString("key")},
	}

	url := viper.GetString("url")
	req, err := http.NewRequest("POST", url, strings.NewReader(parameters.Encode()))
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		var v JResp
		if err := json.Unmarshal(body, &v); err != nil {
			log.Fatal(err)
		}

		if v.Error == 0 {
			log.Print(v.Value)
		} else {
			log.Print(v.ErrorMsg)
		}
	} else {
		log.Print(resp.Status)	
	}
}

// Execute starts stubc execution
func Execute() {
	stubc.Flags().StringP("section", "s", "", "configuration file section")
	stubc.Flags().StringP("key", "k", "", "configuration file key")
	stubc.Flags().StringP("url", "u", "", "url to query")
	viper.BindPFlag("section", stubc.Flags().Lookup("section"))
	viper.BindPFlag("key", stubc.Flags().Lookup("key"))
	viper.BindPFlag("url", stubc.Flags().Lookup("url"))

	if err := stubc.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
