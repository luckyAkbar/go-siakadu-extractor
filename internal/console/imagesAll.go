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
}