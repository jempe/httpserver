package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", "/sdcard", "Define the root filesystem path")

func main() {
	flag.Parse()
	log.Println("Starting web server at http://0.0.0.0:" + *port)

	changeHeader := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			h.ServeHTTP(w, r)
		}
	}

	http.Handle("/", changeHeader(http.FileServer(http.Dir(*root))))
	panic(http.ListenAndServe(":"+*port, nil))
}
