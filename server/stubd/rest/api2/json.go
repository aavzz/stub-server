/*
Package api implements rest api version 2 for stubd
*/
package api2

type JResponse struct {
	Error     int		`json:"error"`
	ErrorMsg  string	`json:"errormsg"`
	Section   string	`json:"section"`
	Key       string	`json:"key"`
	Value  	  string	`json:"value"`
	Signature string	`json:"signature"`
}

type JRequest struct {
	Section   string	`json:"section"`
	Key       string	`json:"key"`
	Signature string	`json:"signature"`
}

