package console

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Use: "Siakadu Information Extractor",
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		log.Print(err.Error())

		os.Exit(-1)
	}
}
