package main

import (
	"io"
	"net/http"
)

func main() {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", "http://localhost:18888", nil)
	req.Proto = "HTTP/1.0"
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(bytes))
}
