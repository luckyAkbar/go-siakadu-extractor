package utility

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func RequestImage(link string) (io.ReadCloser, error) {
	log.Printf("Request image for %s", link)
	res, err := http.Get(link)

	if err != nil {
		log.Panic(err.Error())
	}

	if res.StatusCode != 200 {
		log.Print(fmt.Printf("Failed to get image from %s", link))

		return nil, fmt.Errorf("failed to get image from %s", link)
	}

	return res.Body, nil
}
