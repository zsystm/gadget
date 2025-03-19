package commands

import (
	"github.com/spf13/cobra"

	"github.com/zsystm/gadget/commands/addr"
)

var AddrCmd = &cobra.Command{
	Use:   "addr",
	Short: "addr is a tool for converting between ethereum and bech32 addresses",
}

func init() {
	AddrCmd.AddCommand(addr.BechToEthCmd)
	AddrCmd.AddCommand(addr.EthToBechCmd)
	AddrCmd.AddCommand(addr.ChangeBechPrefixCmd)
}
