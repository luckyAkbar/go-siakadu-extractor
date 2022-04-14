package utility

import (
	"fmt"
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

func (u *Utility) GetMaxSeqFailureCount(cmd *cobra.Command) int {
	maxFailureCount, err := cmd.Flags().GetInt("max-seq-failure")

	if err != nil {
		log.Println("Invalid / no max sequential failure flag present. Using default value 10")

		return 10
	}

	return maxFailureCount
}

func (u *Utility) GetStringFlagValueByName(cmd *cobra.Command, flagName string) string {
	value, err := cmd.Flags().GetString(flagName)

	if err != nil {
		fmt.Printf("Flag value: %s is error %s. Using default value (if any).", flagName, err.Error())
	}

	fmt.Printf("Using flag %s with value %s.\n", flagName, value)

	return value
}

func (u *Utility) GetIntFlagValueByName(cmd *cobra.Command, flagName string) int {
	value, err := cmd.Flags().GetInt(flagName)

	if err != nil {
		fmt.Printf("Flag value: %s is error %s. Using default value (if any).", flagName, err.Error())
	}

	fmt.Printf("Using flag %s with value %d.\n", flagName, value)

	return value
}
