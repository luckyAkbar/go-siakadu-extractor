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

func EnsureNoBackwardNPM(first, second string) error {
	start, err := strconv.Atoi(first)

	if err != nil {
		return fmt.Errorf("NPM value is invalid: %s", first)
	}

	to, err := strconv.Atoi(second)

	if err != nil {
		return fmt.Errorf("NPM value is invalid: %s", first)
	}

	if start > to {
		return fmt.Errorf("backward NPM value detected: %s > %s", first, second)
	}

	if err := ValidateNPM(first); err != nil {
		return fmt.Errorf("invalid value for NPM: %s", first)
	}

	if err := ValidateNPM(second); err != nil {
		return fmt.Errorf("invalid value for NPM: %s", second)
	}

	return nil
}
