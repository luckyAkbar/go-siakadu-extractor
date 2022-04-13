package feature

import (
	"siakadu-extractor/internal/utility"

	"github.com/spf13/cobra"
)

func GetAllProfileImage(cmd *cobra.Command, args []string) {
	util := utility.New()

	delay := util.GetRequestDelaySecFlag(cmd)
	maxSeqFailure := util.GetMaxSeqFailureCount(cmd)
	start := util.GetStringFlagValueByName(cmd, "start")
	to := util.GetStringFlagValueByName(cmd, "to")

	util.GetAllProfileImage(delay, maxSeqFailure, start, to)
}
