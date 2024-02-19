package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gadget",
	Short: "A CLI tool for various tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand. See 'gadget --help' for more details.")
	},
}

var convFromBechCmd = &cobra.Command{
	Use:   "conv-from-bech [bech32-address]",
	Short: "Convert a bech32 address to 20-byte hexadecimal format (= Ethereum address)",
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

func init() {
	rootCmd.AddCommand(convFromBechCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertFromBech32(bech32Address string) (string, error) {
	// Check if the bech32 address is valid
	_, _, err := bech32.Decode(bech32Address)
	if err != nil {
		return "", fmt.Errorf("invalid bech32 address: %v", err)
	}

	// Concatenate "cosmos" to the front of the bech32 address
	result := "cosmos" + bech32Address
	return result, nil
}
