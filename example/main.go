package main

import "./server"

func main() {
	info := server.NewInfo()
	println(info.RedirectURI)
	println(info.ClientID)
	println(info.ClientSecret)
	println(info.StateValue)
//	fb := gofb.NewFacebook(gofb.CLIENT_ID)
//	gofb.Run(8080, fb)
}
