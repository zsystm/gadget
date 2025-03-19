package eth

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var NewAcc = &cobra.Command{
	Use:   "new-acc",
	Short: "Create a random account",
	Args:  cobra.ExactArgs(0),
	RunE:  newAccHandler,
}

func newAccHandler(cmd *cobra.Command, args []string) error {
	key, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	privKey := hexutil.Encode(crypto.FromECDSA(key))
	acc := struct {
		EthPrivKey string         `json:"privateKey"`
		EthAddr    common.Address `json:"address"`
	}{}
	acc.EthPrivKey = privKey
	acc.EthAddr = crypto.PubkeyToAddress(key.PublicKey)
	marshaled, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		return err
	}
	cmd.Printf("%s\n", marshaled)
	return nil
}
