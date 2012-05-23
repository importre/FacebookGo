package gofb

import (
	"./graph"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	ACCESS_INFO_FILE = "access_info.json"
)

type Facebook struct {
	AppID string
	accessInfo

	graph   *graph.Graph
	friends *graph.Friends
}

type accessInfo struct {
	Token   string
	Expires string

	initialized bool
}

func NewFacebook(appid string) *Facebook {
	fb := &Facebook{AppID: appid}
	fb.initialized = false

	data, err := ioutil.ReadFile(ACCESS_INFO_FILE)
	if nil == err {
		json.Unmarshal(data, &fb.accessInfo)
	}

	return fb
}

func (fb *Facebook) Init(params map[string]string) bool {
	if !fb.initialized {
		fb.Token = params["access_token"]
		fb.Expires = params["expires"]
		fb.graph = graph.NewGraph("me", fb.Token)
		fb.friends = graph.NewFriends("me", fb.Token)
		fb.initialized = true
	}
	return true
}

func (fb *Facebook) IsValidSession() bool {
	return false
}

func (fb *Facebook) Initialized() bool {
	return fb.initialized
}

func (fb *Facebook) AccessToken() string {
	return fb.Token
}

func (fb *Facebook) AccessExpires() string {
	return fb.Expires
}

func (fb *Facebook) Graph() *graph.Graph {
  return fb.graph
}

func (fb *Facebook) Friends() *graph.Friends {
  return fb.friends
}

func (fb *Facebook) DumpAccessInfo() {
	jsonData, err := json.Marshal(fb.accessInfo)
	if nil != err {
		log.Println(err)
		return
	}

	buffer := bytes.NewBufferString(string(jsonData))
	ioutil.WriteFile(ACCESS_INFO_FILE, buffer.Bytes(), 0666)
}
