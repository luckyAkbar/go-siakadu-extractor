package console

import (
	"siakadu-extractor/internal/feature"

	"github.com/spf13/cobra"
)

var imageCMD = &cobra.Command{
	Use:     "image",
	Short:   "Get profile image single",
	Long:    "Get profile image from given NPM. If not found, no result produced.",
	Example: "image 1915061056",
	Version: "1",
	Run:     feature.GetImageFromNPM,
}

func init() {
	RootCMD.AddCommand(imageCMD)
}
