package main

import "./server"
import "../gofb"

func main() {
	info := server.InitInfo()
	fb := gofb.NewFacebook(info.ClientID)
	server.Run(8080, fb)
}
