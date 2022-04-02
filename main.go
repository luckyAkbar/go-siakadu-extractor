package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.Get("https://siakadu.unila.ac.id/uploads/fotomhs/thumb/1915061056.jpg")

	if err != nil {
		log.Panic(err.Error())
	}

	defer res.Body.Close()

	file, _ := os.Create("./foto.jpg")

	defer file.Close()

	io.Copy(file, res.Body)
	fmt.Println("beres")
}
