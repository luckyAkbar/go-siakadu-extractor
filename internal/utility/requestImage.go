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
		log.Printf("Failed to request image for %s because: %s", link, err.Error())
		log.Println("Retrying....")
		RequestImage(link)
	}

	if res.StatusCode != 200 {
		log.Print(fmt.Printf("Failed to get image from %s", link))

		return nil, fmt.Errorf("failed to get image from %s", link)
	}

	return res.Body, nil
}
