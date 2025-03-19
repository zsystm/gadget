package cometbft_secp256k1

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var AccInfoFromPrivKeyCmd = &cobra.Command{
	Use:   "acc [prefix] [private-key-hex(no 0x)]",
	Short: "Convert a private hex string(without 0x) to account info",
	Args:  cobra.ExactArgs(2),
	RunE:  accInfoFromPrivateKeyHandler,
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
	cmd.Printf("%s\n", jsonBytes)
	return nil
}
