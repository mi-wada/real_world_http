package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	fmt.Printf("proto: %s\n", resp.Proto)
}
