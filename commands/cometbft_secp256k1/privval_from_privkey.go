package cometbft_secp256k1

import (
	"encoding/hex"
	"os"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/privval"
	"github.com/spf13/cobra"
)

var PrivValFromPrivKeyCmd = &cobra.Command{
	Use:   "privval [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to privval key",
	Args:  cobra.ExactArgs(1),
	RunE:  privValFromPrivateKeyHandler,
}

var privValFromPrivateKeyHandler = func(cmd *cobra.Command, args []string) error {
	var privKeyBytes [secp256k1.PrivKeySize]byte
	input, err := hex.DecodeString(args[0])
	if err != nil {
		return err
	}
	copy(privKeyBytes[:], input)
	privKey := secp256k1.PrivKey(privKeyBytes[:])
	// get current directory path
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	pv := privval.NewFilePV(privKey, dir+"privval", dir+"state")
	pv.Save()

	jsonBytes, err := cmtjson.MarshalIndent(pv, "", "  ")
	if err != nil {
		panic(err)
	}

	cmd.Printf("privVal:%s\n", jsonBytes)
	return nil
}
