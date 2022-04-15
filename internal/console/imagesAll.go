package console

import (
	"siakadu-extractor/internal/feature"

	"github.com/spf13/cobra"
)

var getAllImagesCmd = &cobra.Command{
	Use:     "all",
	Short:   "Get all public profile picture",
	Long:    "Get all public profile picture from all posible NPM",
	Example: "images all",
	Version: "1",
	Run:     feature.GetAllProfileImage,
}

func init() {
	imagesCmd.AddCommand(getAllImagesCmd)
	getAllImagesCmd.PersistentFlags().Int("request-delay-sec", 0, "Make every request delayed as input (ms)")
	getAllImagesCmd.PersistentFlags().Int("use-image-req-optimizer", 10, "skip and increment division if sequentially failed to get images from given NPM when sequentially absent crawling is failed / not found. Default to 10")
	getAllImagesCmd.PersistentFlags().String("start", "0000000000", "Set starting point for full images crawling. Default to: '0000000000'")
	getAllImagesCmd.PersistentFlags().String("to", "9999999999", "Set final point for full images crawling. Default to: '9999999999'")
}
