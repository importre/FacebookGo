package main

import "./gofb"

func main() {
	fb := gofb.NewFacebook(gofb.CLIENT_ID)
	gofb.Run(8080, fb)
}
