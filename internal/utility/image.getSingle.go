package utility

import (
	"fmt"
	"log"
)

func (u *Utility) GetImageFromNPM(NPM string) error {
	for _, ext := range POSSIBLE_IMAGE_EXT {
		link := fmt.Sprintf("%s%s%s", IMAGE_URL, NPM, ext)
		res, err := RequestImage(link)
		if err == nil {
			saveImage(res, NPM, ext)
			u.WriteLog(fmt.Sprintf("Image for NPM: %s found.", NPM))

			return nil
		}
		defer handlePanic(link)
	}

	u.WriteLog(fmt.Sprintf("Image for NPM: %s is not found.", NPM))

	return fmt.Errorf("Image for NPM: %s is not found.", NPM)
}

func handlePanic(link string) {
	if r := recover(); r != nil {
		log.Println("WARNING panic occured:", r)
		log.Println("Recovering from panic when request image from:", link)
		log.Println("Program should continue running...")
	}
}
