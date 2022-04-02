package utility

import (
	"fmt"
	"strconv"
)

func ValidateNPM(NPM string) error {
	if len(NPM) != 10 {
		return fmt.Errorf("invalid value in argument NPM: %s is not a valid NPM", NPM)
	}

	_, err := strconv.Atoi(NPM)

	if err != nil {
		return fmt.Errorf("invalid value in argument NPM: NPM must be a valid numeric string\n")
	}

	return nil
}
