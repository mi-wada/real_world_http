package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func fetch_all() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func fetch_seq() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}

func fetch_seq_socket() {
	// Prepare TCP connection
	dialer := &net.Dialer{}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Request
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:18888", nil)
	// Send request
	err = req.Write(conn)
	if err != nil {
		panic(err)
	}

	// Response
	reader := bufio.NewReader(conn)
	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}
	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		line := make([]byte, size)
		reader.Read(line)
		reader.Discard(2)
		log.Printf("  %s\n", strings.TrimSpace(string(line)))
	}
}

func main() {
	// fetch_seq_socket()

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:18888", nil)
	// req.Write to string
	var buf bytes.Buffer
	req.Write(&buf)
	fmt.Println(buf.String())
}
