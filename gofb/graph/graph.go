package graph

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_GRAPH_URL = "https://graph.facebook.com/%v"
)

type Grapher interface {
	Query() string
}

type Graph struct {
	baseUrl string
}

func NewGraph(uid string) (graph *Graph) {
	baseUrl = fmt.Sprintf(BASE_GRAPH_URL, uid)
	graph := &Graph{baseUrl}
	return
}

func NewMe() (graph *Graph) {
	baseUrl = fmt.Sprintf(BASE_GRAPH_URL, "me")
	graph := &Graph{baseUrl}
	return
}

func (g *Graph) Init() bool {
	//	url := "https://graph.facebook.com/me?access_token=%v"
	//	url = fmt.Sprintf(url, fb.Token)
	//	resp, _ := http.Get(url)
	//	defer resp.Body.Close()
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println(string(body))
	return true
}

func (g *Graph) Query() string {
	return ""
}

func (g *Graph) RawJson(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonData := string(body)
	return jsonData
}
