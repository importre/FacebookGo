package server

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

const (
	INFO_FILE = "info.json"
)

type Info struct {
	RedirectURI  string
	ClientID     string
	ClientSecret string
	StateValue   string
}

func NewInfo() *Info {
	info := &Info{}
	data, err := ioutil.ReadFile(INFO_FILE)
	if nil != err {
		log.Fatal(err)
	}
	json.Unmarshal(data, info)
	return info
}
