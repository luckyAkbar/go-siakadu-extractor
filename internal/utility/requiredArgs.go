package utility

import "fmt"

func RequiredArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing required command arguments")
	}

	return nil
}

func RequiredArgsNum(args []string, num int) error {
	if len(args) != num {
		return fmt.Errorf("command need exactly %d argument(s). Got: %d", num, len(args))
	}

	return nil
}
