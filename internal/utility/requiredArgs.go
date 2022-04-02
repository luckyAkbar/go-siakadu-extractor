package utility

import "fmt"

func RequiredArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing required command arguments")
	}

	return nil
}
