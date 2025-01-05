package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello via %s", r.Proto)
	})

	log.Println("start at https://localhost:8443")
	log.Println(http3.ListenAndServeQUIC("0.0.0.0:8443", "localhost.pem", "localhost-key.pem", mux))
}
