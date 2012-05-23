package graph

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Friend struct {
	Name, Id string
}

type Friends struct {
	baseUrl     string
	accessToken string
	Data []Friend
}

func NewFriends(uid, accessToken string) (friends *Friends) {
	baseUrl := fmt.Sprintf(BASE_GRAPH_URL, uid+"/friends")
	friends = &Friends{baseUrl: baseUrl, accessToken: accessToken}
	friends.Query()
	return
}

func (f *Friends) Query() {
	var url string
	if "" != f.accessToken {
		url = "?access_token=" + f.accessToken
	}

	url = f.baseUrl + url
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, f)
	fmt.Println(err)
	//for i, j := range(f.Data) {
	//	fmt.Println(i, j)
	//}
}
