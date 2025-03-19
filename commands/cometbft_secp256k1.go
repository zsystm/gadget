package commands

import (
	"github.com/spf13/cobra"

	"github.com/zsystm/gadget/commands/cometbft_secp256k1"
)

var Secp256k1Cmd = &cobra.Command{
	Use:   "secp256k1",
	Short: "secp256k1 is a tool for secp256k1 related operations",
}

func init() {
	Secp256k1Cmd.AddCommand(cometbft_secp256k1.PubKeyFromPrivKeyCmd)
	Secp256k1Cmd.AddCommand(cometbft_secp256k1.PrivValFromPrivKeyCmd)
	Secp256k1Cmd.AddCommand(cometbft_secp256k1.AccInfoFromPrivKeyCmd)
}
