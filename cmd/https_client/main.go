package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	url      = "https://localhost:18443"
	certFile = "./cmd/https_server/ca.crt"
)

func main() {
	// 証明書を読み込む
	cert, err := os.ReadFile(certFile)
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	// クライアントを作成
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// 通信を行う
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
