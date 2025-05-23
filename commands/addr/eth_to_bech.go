package addr

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/spf13/cobra"
)

var EthToBechCmd = &cobra.Command{
	Use:   "eth-to-bech [eth-address] [prefix]",
	Short: "Convert a 20-byte hexadecimal ethereum address to bech32 address",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ethAddress := args[0]
		// if ethAddress has 0x prefix, remove it
		if ethAddress[:2] == "0x" {
			ethAddress = ethAddress[2:]
		}
		decoded, err := hex.DecodeString(ethAddress)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		converted, err := bech32.ConvertBits(decoded, 8, 5, true)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		prefix := args[1]
		bech32Address, err := bech32.Encode(prefix, converted)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		fmt.Printf("converted: %s\n", bech32Address)
		return nil
	},
}
