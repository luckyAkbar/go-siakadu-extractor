package utility

import (
	"fmt"
)

func (u *Utility) GetImageFromNPM(NPM string) error {
	for _, ext := range POSSIBLE_IMAGE_EXT {
		res, err := RequestImage(fmt.Sprintf("%s%s%s", IMAGE_URL, NPM, ext))

		if err == nil {
			saveImage(res, NPM, ext)
			res.Close()

			u.WriteLog(fmt.Sprintf("Image for NPM: %s found.", NPM))

			return nil
		}
	}

	u.WriteLog(fmt.Sprintf("Image for NPM: %s is not found.", NPM))

	return fmt.Errorf("Image for NPM: %s is not found.", NPM)
}
