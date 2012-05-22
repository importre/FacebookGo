package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Initializer interface {
	Init(map[string]string) bool
}

const (
	DIALOG_BASE_URL  = "https://www.facebook.com/dialog/oauth?"
	ACCESS_TOKEN_URL = "https://graph.facebook.com/oauth/access_token?"
)

var (
	SERVER_URI    = "http://127.0.0.1"
	REDIRECT_URI  = SERVER_URI + ":%v/auth/"
	CLIENT_ID     = "295362970552833"
	CLIENT_SECRET = "9752635279a076efc9cc31c11a138f1e"
	STATE_VALUE   = "HelloFacebook"

	initializer Initializer
)

func init() {
	http.Handle("/", http.HandlerFunc(MainHandler))
	http.Handle("/auth/", http.HandlerFunc(AuthHandler))
}

func Run(port uint, init Initializer) {
	REDIRECT_URI = fmt.Sprintf(REDIRECT_URI, port)
	initializer = init

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	if nil != err {
		log.Fatal("ListenAndServe:", err)
	}
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	state, code := params.Get("state"), params.Get("code")

	if STATE_VALUE != state || "" == code {
		// error
		log.Println(1)
		return
	}

	redirectUrl := accessTokenUrl(code)
	resp, err := http.Get(redirectUrl)
	if nil != err {
		// error
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		// error
		log.Println(err)
		return
	}

	params, err = url.ParseQuery(string(body))
	if nil != err {
		// error
		log.Println(err)
		return
	}

	accessToken := params.Get("access_token")
	expires := params.Get("expires")
	if "" == accessToken || "" == expires {
		// error
		log.Println(2)
		return
	}

	initParams := map[string]string{
		"access_token": accessToken,
		"expires":      expires,
	}

	initializer.Init(initParams)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	redirectUrl := loginUrlString()
	http.Redirect(w, r, redirectUrl, http.StatusFound)
}

func urlString(baseUrl string, params *url.Values) string {
	return baseUrl + params.Encode()
}

func loginUrlString() string {
	v := &url.Values{}
	v.Set("client_id", CLIENT_ID)
	v.Set("redirect_uri", REDIRECT_URI)
	v.Set("state", STATE_VALUE)
	return urlString(DIALOG_BASE_URL, v)
}

func accessTokenUrl(code string) string {
	v := &url.Values{}
	v.Set("client_id", CLIENT_ID)
	v.Set("redirect_uri", REDIRECT_URI)
	v.Set("client_secret", CLIENT_SECRET)
	v.Set("code", code)
	return urlString(ACCESS_TOKEN_URL, v)
}
