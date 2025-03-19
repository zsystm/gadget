package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var Base64ToHexCmd = &cobra.Command{
	Use:   "b64-to-hex [b64-string]",
	Short: "Convert a base64 string to hexadecimal",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		b64 := args[0]
		decoded, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			return err
		}
		fmt.Printf("hex: %s\n", common.Bytes2Hex(decoded))
		return nil
	},
}

var HexToBase64Cmd = &cobra.Command{
	Use:   "hex-to-b64 [hex-string]",
	Short: "Convert a hexadecimal string to base64",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hexStr := args[0]
		bytes := common.Hex2Bytes(hexStr)
		encoded := base64.StdEncoding.EncodeToString(bytes)
		fmt.Printf("b64: %s\n", encoded)
		return nil
	},
}
