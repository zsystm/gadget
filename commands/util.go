package commands

import (
	"github.com/spf13/cobra"

	"github.com/zsystm/gadget/commands/util"
)

var UtilCmd = &cobra.Command{
	Use:   "util",
	Short: "util provides small helper utilities",
}

func init() {
	UtilCmd.AddCommand(util.TimestampCmd)
}
