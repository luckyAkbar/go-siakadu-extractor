package feature

import (
	"fmt"
	"os"
	"siakadu-extractor/internal/utility"

	"github.com/spf13/cobra"
)

func GetImageFromNPM(cmd *cobra.Command, args []string) {
	util := utility.New()

	if err := utility.RequiredArgs(args); err != nil {
		fmt.Println(err.Error())

		os.Exit(-1)
	}

	if err := utility.ValidateNPM(args[0]); err != nil {
		fmt.Printf(err.Error())

		os.Exit(-1)
	}

	util.GetImageFromNPM(args[0])
}
