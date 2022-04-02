package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"siakadu-extractor/internal/console"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	if _, err := os.Stat("./result"); os.IsNotExist(err) {
		os.MkdirAll("result/images", os.ModePerm)
	}

	if _, err := os.Stat("./report"); os.IsNotExist(err) {
		os.Mkdir("report", os.ModePerm)
	}

	console.Execute()
}
