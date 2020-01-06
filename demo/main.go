package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	dir := "D:/gohub/src/im/demo"
	//dir := "./"
	log.Fatal(http.ListenAndServe(":1999", http.FileServer(http.Dir(dir))))
	// log.Fatal(http.ListenAndServeTLS(":1999", "/etc/letsencrypt/live/texixi.com-0002/fullchain.pem", "/etc/letsencrypt/live/texixi.com-0002/privkey.pem", http.FileServer(http.Dir("./"))))

}
