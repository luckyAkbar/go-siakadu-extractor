package feature

import (
	"siakadu-extractor/internal/utility"

	"github.com/spf13/cobra"
)

func GetAllProfileImage(cmd *cobra.Command, args []string) {
	util := utility.New()

	delay := util.GetIntFlagValueByName(cmd, "request-delay-sec")
	maxSeqFailure := util.GetIntFlagValueByName(cmd, "use-image-req-optimizer")
	start := util.GetStringFlagValueByName(cmd, "start")
	to := util.GetStringFlagValueByName(cmd, "to")
	skipYearPlus1 := util.GetBoolFLagValueByName(cmd, "skip-year-plus-1")

	util.GetAllProfileImage(delay, maxSeqFailure, start, to, skipYearPlus1)
}
