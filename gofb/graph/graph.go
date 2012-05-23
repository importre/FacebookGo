package graph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_GRAPH_URL = "https://graph.facebook.com/%v"
)

type Grapher interface {
	Query(accessToken string)
}

type Graph struct {
	baseUrl               string
	accessToken           string
	Id, Name, Username    string
	First_name, Last_name string
	Gender                string
}

func NewGraph(uid, accessToken string) (graph *Graph) {
	baseUrl := fmt.Sprintf(BASE_GRAPH_URL, uid)
	graph = &Graph{baseUrl: baseUrl, accessToken: accessToken}
	graph.Query()
	return
}

func (g *Graph) Query() {
	var url string
	if "" != g.accessToken {
		url = "?access_token=" + g.accessToken
	}

	url = g.baseUrl + url
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, g)
	fmt.Println(err)
	//fmt.Println(*g)
}
