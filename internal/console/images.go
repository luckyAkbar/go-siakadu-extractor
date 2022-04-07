package console

import (
	"siakadu-extractor/internal/feature"

	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:     "images",
	Short:   "Get profile image range",
	Long:    "Get profile image from given NPM range. If not found, no result produced.",
	Example: "image 1915061056 1915061060",
	Version: "1",
	Run:     feature.GetImageFromNPMRange,
}

func init() {
	RootCMD.AddCommand(imagesCmd)
	imagesCmd.PersistentFlags().Int("request-delay-sec", 0, "Make every request delayed as input (ms)")
}
