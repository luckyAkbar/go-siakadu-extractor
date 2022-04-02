package utility

import (
	"fmt"
	"io"
	"log"
	"os"
)

func saveImage(body io.ReadCloser, NPM string, ext string) {
	log.Printf("saving image file for: %s", NPM)

	file, _ := os.Create(fmt.Sprintf("%s%s%s", BASE_IMAGE_STORAGE, NPM, ext)) // fmt.Sprintf("%s%s%s", BASE_IMAGE_STORAGE, NPM, ext)

	_, err := io.Copy(file, body)

	if err != nil {
		log.Print(err)
	}

	defer file.Close()

	log.Printf("Success saving image for: %s", NPM)
}
