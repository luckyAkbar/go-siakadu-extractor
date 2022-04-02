package utility

import (
	"fmt"
	"log"
	"os"
)

func (u *Utility) GetImageFromNPM(NPM string) {
	for _, ext := range POSSIBLE_IMAGE_EXT {
		res, err := RequestImage(fmt.Sprintf("%s%s%s", IMAGE_URL, NPM, ext))

		if err == nil {
			saveImage(res, NPM, ext)

			res.Close()

			os.Exit(0)
		}
	}

	log.Printf("Image for NPM: %s is not found.", NPM)
}
