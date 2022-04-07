package console

import (
	"siakadu-extractor/internal/feature"

	"github.com/spf13/cobra"
)

var getAllImagesCmd = &cobra.Command{
	Use: "all",
	Short: "Get all public profile picture",
	Long: "Get all public profile picture from all posible NPM",
	Example: "images all",
	Version: "1",
	Run: feature.GetAllProfileImage,
}

func init() {
	imagesCmd.AddCommand(getAllImagesCmd)
	getAllImagesCmd.PersistentFlags().Int("request-delay-sec", 0, "Make every request delayed as input (ms)")
	getAllImagesCmd.PersistentFlags().Int("max-seq-failure", 10, "skip and increment division if sequentially failed to get images from given NPM when sequentially absent crawling is failed / not found")
}