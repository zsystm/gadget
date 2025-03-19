package cometbft_ed25519

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var PubKeyFromPrivKeyCmd = &cobra.Command{
	Use:   "pubkey-from-privkey [privkey]",
	Short: "Generate public key from private key",
	Args:  cobra.ExactArgs(1),
	RunE:  pubKeyFromPrivateKeyHandler,
}

var pubKeyFromPrivateKeyHandler = func(cmd *cobra.Command, args []string) error {
	var privateKeyBytes [ed25519.PrivateKeySize]byte
	input, err := hex.DecodeString(args[0])
	if err != nil {
		return err
	}
	copy(privateKeyBytes[:], input)
	privateKey := ed25519.PrivKey(privateKeyBytes[:])
	publicKey := privateKey.PubKey()
	hex := common.Bytes2Hex(privateKey[:])
	b64 := base64.StdEncoding.EncodeToString(publicKey.Bytes())
	cmd.Printf("public key: {hex: %s, base64: %s}\n", hex, b64)
	return nil
}
