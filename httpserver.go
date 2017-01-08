package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", "/sdcard", "Define the root filesystem path")

func main() {
	flag.Parse()

	changeHeader := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			h.ServeHTTP(w, r)
		}
	}

	http.Handle("/", changeHeader(http.FileServer(http.Dir(*root))))

	localIP := getLocalIP()

	fmt.Println("Starting web server at http://" + localIP + ":" + *port)

	panic(http.ListenAndServe(":"+*port, nil))
}

// GetLocalIP returns the non loopback local IP of the host
func getLocalIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		ip := address.(*net.IPNet)
		if ip.IP.To4() != nil && !ip.IP.IsLoopback() {
			return ip.IP.String()
		}
	}
	return ""
}
