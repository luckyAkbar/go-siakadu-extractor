package feature

import (
	"log"
	"os"
	"siakadu-extractor/internal/utility"

	"github.com/spf13/cobra"
)

func GetImageFromNPMRange(cmd *cobra.Command, args []string) {
	util := utility.New()
	delay := util.GetRequestDelaySecFlag(cmd)

	if err := utility.RequiredArgsNum(args, 2); err != nil {
		log.Print(err.Error())

		os.Exit(-1)
	}

	for _, NPM := range args {
		if err := utility.ValidateNPM(NPM); err != nil {
			log.Println(err.Error())

			os.Exit(-1)
		}
	}

	util.GetImageFromNPMRange(args[0], args[1], delay)
	os.Exit(0)
}
