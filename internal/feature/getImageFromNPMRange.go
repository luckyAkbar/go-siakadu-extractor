package feature

import (
	"log"
	"os"
	"siakadu-extractor/internal/utility"

	"github.com/spf13/cobra"
)

func GetImageFromNPMRange(cmd *cobra.Command, args []string) {
	var requestOptimizer int = 0
	util := utility.New()
	delay := util.GetRequestDelaySecFlag(cmd)

	if val := util.GetIntFlagValueByName(cmd, "use-image-req-optimizer"); val != 0 {
		log.Printf("Image request optimizer enabled with value: %d\n", val)

		requestOptimizer = val
	}

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

	util.GetImageFromNPMRange(args[0], args[1], delay, requestOptimizer)
	os.Exit(0)
}
