package utility

import "fmt"

func generatePossibleImageName(NPM string) []string {
	var links []string

	for _, ext := range POSSIBLE_IMAGE_EXT {
		links = append(links, fmt.Sprintf("%s%s%s", IMAGE_URL, NPM, ext))
	}

	return links
}
