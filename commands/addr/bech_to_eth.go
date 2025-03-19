package addr

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/spf13/cobra"
)

var BechToEthCmd = &cobra.Command{
	Use:   "bech-to-eth [bech32-address]",
	Short: "Convert a bech32 address to 20-byte hexadecimal ethereum address)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bech32Address := args[0]
		_, decoded, err := bech32.Decode(bech32Address)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		converted, err := bech32.ConvertBits(decoded, 5, 8, false)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		fmt.Printf("converted: 0x%s\n", hex.EncodeToString(converted))
		return nil
	},
}
