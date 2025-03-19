package addr

import (
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/spf13/cobra"
)

var ChangeBechPrefixCmd = &cobra.Command{
	Use:   "change-bech-prefix [bech32-address] [prefix]",
	Short: "Convert a bech32 address to another prefix",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bech32Address := args[0]
		_, decoded, err := bech32.Decode(bech32Address)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		bytes, err := bech32.ConvertBits(decoded, 5, 8, false)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		prefix := args[1]
		converted, err := bech32.ConvertBits(bytes, 8, 5, true)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		newBech32Address, err := bech32.Encode(prefix, converted)
		if err != nil {
			_ = fmt.Errorf("Error: %v\n", err)
			return err
		}
		fmt.Printf("converted: %s\n", newBech32Address)
		return nil
	},
}
