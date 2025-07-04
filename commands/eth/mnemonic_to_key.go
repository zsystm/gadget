package eth

import (
	"encoding/json"
	"strings"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
	"github.com/zsystm/gadget/pkg/types"
)

var MnemonicToKeyCmd = &cobra.Command{
	Use:   "mnemonic [words]",
	Short: "Derive an Ethereum account from a BIP-39 mnemonic",
	Args:  cobra.MinimumNArgs(1),
	RunE:  mnemonicToKeyHandler,
}

func mnemonicToKeyHandler(cmd *cobra.Command, args []string) error {
	mnemonic := strings.Join(args, " ")
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return err
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return err
	}
	privKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return err
	}
	acc := types.EthAcc{
		EthPrivKey: privKeyHex,
		EthAddr:    account.Address.Hex(),
	}
	out, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		return err
	}
	cmd.Printf("%s\n", out)
	return nil
}
