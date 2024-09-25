package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var pubKeyFromPrivateKeyCmd = &cobra.Command{
	Use:   "secp256k1-pub [private-key-hex]",
	Short: "Convert a base64 string to hexadecimal",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var privKeyBytes [secp256k1.PrivKeySize]byte
		input, err := hex.DecodeString(args[0])
		if err != nil {
			return err
		}
		copy(privKeyBytes[:], input)
		privKey := secp256k1.PrivKey(privKeyBytes[:])
		pubKey := privKey.PubKey()
		hex := common.Bytes2Hex(pubKey.Bytes())
		b64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())
		fmt.Printf("pubKey:{hex: %s, base64: %s}\n", hex, b64)
		return nil
	},
}
