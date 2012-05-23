package server

import (
	"../../gofb/graph"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Initializer interface {
	Init(map[string]string) bool
	Graph() *graph.Graph
	Friends() *graph.Friends
}

type Info struct {
	ServerURI    string
	ClientID     string
	ClientSecret string
	StateValue   string
}

const (
	DIALOG_BASE_URL   = "https://www.facebook.com/dialog/oauth?"
	ACCESS_TOKEN_URL  = "https://graph.facebook.com/oauth/access_token?"
	INFO_FILE         = "info.json"
	USER_TMPL_FILE    = "user.html"
	FRIENDS_TMPL_FILE = "friends.html"
	AUTH_PATH         = "/auth/"
	USER_PATH         = "/user/"
	FRIENDS_PATH      = "/friends/"
)

var (
	initializer Initializer
	info        *Info
	redirectUri string
)

func init() {
	http.Handle("/", http.HandlerFunc(MainHandler))
	http.Handle(AUTH_PATH, http.HandlerFunc(AuthHandler))
	http.Handle(USER_PATH, http.HandlerFunc(UserHandler))
	http.Handle(FRIENDS_PATH, http.HandlerFunc(FriendsHandler))
}

func InitInfo() *Info {
	info = &Info{}
	data, err := ioutil.ReadFile(INFO_FILE)
	if nil != err {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, info)
	if nil != err {
		log.Fatal(err)
	}

	info.ServerURI += ":%v"
	redirectUri = info.ServerURI + AUTH_PATH
	return info
}

func Run(port uint, init Initializer) {
	info.ServerURI = fmt.Sprintf(info.ServerURI, port)
	redirectUri = fmt.Sprintf(redirectUri, port)
	initializer = init

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	if nil != err {
		log.Fatal("ListenAndServe:", err)
	}
}

func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New(FRIENDS_TMPL_FILE)
	t, _ = t.ParseFiles(FRIENDS_TMPL_FILE)
	t.Execute(w, initializer.Friends().Data)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New(USER_TMPL_FILE)
	t, _ = t.ParseFiles(USER_TMPL_FILE)
	t.Execute(w, initializer.Graph())
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	state, code := params.Get("state"), params.Get("code")

	if info.StateValue != state || "" == code {
		log.Println(1)
		return
	}

	redirectUrl := accessTokenUrl(code)
	resp, err := http.Get(redirectUrl)
	if nil != err {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return
	}

	params, err = url.ParseQuery(string(body))
	if nil != err {
		log.Println(err)
		return
	}

	accessToken := params.Get("access_token")
	expires := params.Get("expires")
	if "" == accessToken || "" == expires {
		log.Println(2)
		return
	}

	initParams := map[string]string{
		"access_token": accessToken,
		"expires":      expires,
	}

	initializer.Init(initParams)
	http.Redirect(w, r, USER_PATH, http.StatusMovedPermanently)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	redirectUri := loginUrlString()
	http.Redirect(w, r, redirectUri, http.StatusFound)
}

func urlString(baseUrl string, params *url.Values) string {
	return baseUrl + params.Encode()
}

func loginUrlString() string {
	v := &url.Values{}
	v.Set("client_id", info.ClientID)
	v.Set("redirect_uri", redirectUri)
	v.Set("state", info.StateValue)
	return urlString(DIALOG_BASE_URL, v)
}

func accessTokenUrl(code string) string {
	v := &url.Values{}
	v.Set("client_id", info.ClientID)
	v.Set("redirect_uri", redirectUri)
	v.Set("client_secret", info.ClientSecret)
	v.Set("code", code)
	return urlString(ACCESS_TOKEN_URL, v)
}
