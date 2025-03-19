package commands

import (
	"github.com/spf13/cobra"

	"github.com/zsystm/gadget/commands/cometbft_ed25519"
)

var Ed25519Cmd = &cobra.Command{
	Use:   "ed25519",
	Short: "ed25519 is a tool for ed25519 key operations",
}

func init() {
	Ed25519Cmd.AddCommand(cometbft_ed25519.PubKeyFromPrivKeyCmd)
	Ed25519Cmd.AddCommand(cometbft_ed25519.AddrFromPubKeyCmd)
}
