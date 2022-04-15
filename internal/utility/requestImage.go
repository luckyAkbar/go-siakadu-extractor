package utility

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func RequestImage(link string) (io.ReadCloser, error) {
	log.Println(fmt.Sprintf("Request image for %s.", link))
	res, err := http.Get(link)

	if err != nil {
		log.Println(fmt.Sprintf("Failed to request image for %s because: %s", link, err.Error()))
		log.Println("Retrying....")
		RequestImage(link)
	}

	if res.StatusCode != 200 {
		log.Println("Failed to get image from: ", link)

		return nil, fmt.Errorf("failed to get image from %s", link)
	}

	return res.Body, nil
}
