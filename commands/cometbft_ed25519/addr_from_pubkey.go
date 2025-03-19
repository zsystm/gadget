package cometbft_ed25519

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/spf13/cobra"
)

var AddrFromPubKeyCmd = &cobra.Command{
	Use:   "addr-from-pubkey",
	Short: "Generate address from public key",
	Args:  cobra.ExactArgs(1),
	RunE:  addrFromPubKeyHandler,
}

var addrFromPubKeyHandler = func(cmd *cobra.Command, args []string) error {
	var decoded []byte
	var err error

	// first try to decode the input as base64
	input := args[0]
	decoded, err = base64.StdEncoding.DecodeString(input)
	if err != nil {
		// try to decode the input as hex
		decoded, err = hex.DecodeString(input)
		if err != nil {
			return err
		}
	}

	pubKey := ed25519.PubKey(decoded)
	addr := pubKey.Address()

	cmd.Printf("address: %s\n", addr)
	return nil
}
