package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	port     = 18443
	certFile = "./cmd/https_server/server.crt"
	keyFile  = "./cmd/https_server/server.key"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("start http listening :%d\n", port)
	log.Println(http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil))
}
