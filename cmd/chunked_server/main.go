package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = 18888

func handler(w http.ResponseWriter, r *http.Request) {
	flusher := w.(http.Flusher)
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "chunk %d\n", i)
		flusher.Flush()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Printf("start http listening :%d\n", port)
	httpServer.Addr = fmt.Sprintf(":%d", port)
	log.Println(httpServer.ListenAndServe())
}
