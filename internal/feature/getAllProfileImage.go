package feature

import (
	"siakadu-extractor/internal/utility"

	"github.com/spf13/cobra"
)

func GetAllProfileImage(cmd *cobra.Command, args []string) {
	util := utility.New()

	util.GetAllProfileImage()
}