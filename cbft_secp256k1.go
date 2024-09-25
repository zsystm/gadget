package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/privval"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var secp256k1Cmd = &cobra.Command{
	Use:   "secp256k1",
	Short: "secp256k1 related commands",
	Long:  "All commands related to cometbft secp256k1",
}

var pubKeyFromPrivateKeyCmd = &cobra.Command{
	Use:   "pub [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to public key",
	Args:  cobra.ExactArgs(1),
	RunE:  pubKeyFromPrivateKeyHandler,
}

var privValFromPrivateKeyCmd = &cobra.Command{
	Use:   "privval [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to privval key",
	Args:  cobra.ExactArgs(1),
	RunE:  privValFromPrivateKeyHandler,
}

var genesisInfoFromPrivateKeyCmd = &cobra.Command{
	Use:   "genesis [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to genesis info",
	Args:  cobra.ExactArgs(1),
	RunE:  genesisInfoFromPrivateKeyHandler,
}

func init() {
	secp256k1Cmd.AddCommand(pubKeyFromPrivateKeyCmd)
	secp256k1Cmd.AddCommand(privValFromPrivateKeyCmd)
	secp256k1Cmd.AddCommand(genesisInfoFromPrivateKeyCmd)
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
	fmt.Printf("pubKey:{hex: %s, base64: %s}\n", hex, b64)
	return nil
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

	fmt.Printf("privVal:%s\n", jsonBytes)
	return nil
}

var genesisInfoFromPrivateKeyHandler = func(cmd *cobra.Command, args []string) error {
	var privKeyBytes [secp256k1.PrivKeySize]byte
	input, err := hex.DecodeString(args[0])
	if err != nil {
		return err
	}
	copy(privKeyBytes[:], input)
	privKey := secp256k1.PrivKey(privKeyBytes[:])
	pubKey := privKey.PubKey()

	accAddr := sdk.AccAddress(pubKey.Address().Bytes()).String()
	valAddr := sdk.ValAddress(pubKey.Address().Bytes()).String()
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())

	output := struct {
		AccAddr string `json:"accAddr"`
		ValAddr string `json:"valAddr"`
		PubKey  string `json:"pubKey"`
	}{
		AccAddr: accAddr,
		ValAddr: valAddr,
		PubKey:  pubKeyBase64,
	}

	jsonBytes, err := cmtjson.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("genesisInfo:%s\n", jsonBytes)
	return nil
}
