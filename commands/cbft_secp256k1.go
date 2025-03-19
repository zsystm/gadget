package commands

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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var Secp256k1Cmd = &cobra.Command{
	Use:   "secp256k1",
	Short: "secp256k1 related commands",
	Long:  "All commands related to cometbft secp256k1",
}

var PubKeyFromPrivateKeyCmd = &cobra.Command{
	Use:   "pub [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to public key",
	Args:  cobra.ExactArgs(1),
	RunE:  pubKeyFromPrivateKeyHandler,
}

var PrivValFromPrivateKeyCmd = &cobra.Command{
	Use:   "privval [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to privval key",
	Args:  cobra.ExactArgs(1),
	RunE:  privValFromPrivateKeyHandler,
}

var AccInfoFromPrivateKeyCmd = &cobra.Command{
	Use:   "acc [prefix] [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to account info",
	Args:  cobra.ExactArgs(2),
	RunE:  accInfoFromPrivateKeyHandler,
}

func init() {
	Secp256k1Cmd.AddCommand(PubKeyFromPrivateKeyCmd)
	Secp256k1Cmd.AddCommand(PrivValFromPrivateKeyCmd)
	Secp256k1Cmd.AddCommand(AccInfoFromPrivateKeyCmd)
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

var accInfoFromPrivateKeyHandler = func(cmd *cobra.Command, args []string) error {
	prefix := args[0]
	inputPrivKey := args[1]
	ethPrivKey, err := crypto.HexToECDSA(inputPrivKey)
	if err != nil {
		return err
	}
	ethAddr := crypto.PubkeyToAddress(ethPrivKey.PublicKey).Hex()
	input, err := hex.DecodeString(inputPrivKey)
	if err != nil {
		return err
	}
	var privKeyBytes [secp256k1.PrivKeySize]byte
	copy(privKeyBytes[:], input)
	privKey := secp256k1.PrivKey(privKeyBytes[:])
	pubKey := privKey.PubKey()

	conf := sdk.GetConfig()
	conf.SetBech32PrefixForAccount(prefix, prefix+"pub")
	conf.SetBech32PrefixForValidator(prefix+"valoper", prefix+"valoperpub")
	conf.SetBech32PrefixForConsensusNode(prefix+"valcons", prefix+"valconspub")
	accAddr := sdk.AccAddress(pubKey.Address().Bytes()).String()
	valAddr := sdk.ValAddress(pubKey.Address().Bytes()).String()
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey.Bytes())

	output := struct {
		EthAddr string `json:"ethAddr"`
		AccAddr string `json:"accAddr"`
		ValAddr string `json:"valAddr"`
		PubKey  string `json:"pubKey"`
	}{
		EthAddr: ethAddr,
		AccAddr: accAddr,
		ValAddr: valAddr,
		PubKey:  pubKeyBase64,
	}

	jsonBytes, err := cmtjson.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonBytes)
	return nil
}
