package cometbft_secp256k1

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var PubKeyFromPrivKeyCmd = &cobra.Command{
	Use:   "pub [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to public key",
	Args:  cobra.ExactArgs(1),
	RunE:  pubKeyFromPrivateKeyHandler,
}

var pubKeyFromPrivateKeyHandler = func(cmd *cobra.Command, args []string) error {
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
	cmd.Printf("pubKey:{hex: %s, base64: %s}\n", hex, b64)
	return nil
}
