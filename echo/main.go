package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const port = 18888

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Printf("r.Header.Get(\"X-Test\"): %v\n", r.Header.Get("X-Test"))
	fmt.Printf("r.Header.Values(\"X-Test\"): %v\n", r.Header.Values("X-Test"))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Printf("start http listening :%d\n", port)
	httpServer.Addr = fmt.Sprintf(":%d", port)
	log.Println(httpServer.ListenAndServe())
}
