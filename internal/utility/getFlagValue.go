package utility

import (
	"log"

	"github.com/spf13/cobra"
)

func (u *Utility) GetRequestDelaySecFlag(cmd *cobra.Command) int {
	delay, err := cmd.Flags().GetInt("request-delay-sec")

	if err != nil {
		log.Println("Invalid / no request delay present. Using default value 0s")

		return 0
	}

	return delay
}