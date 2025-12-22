package eth

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	cmdhex "github.com/zsystm/gadget/commands/internal/hexutil"
)

var GetAddressFromPrivateKey = &cobra.Command{
	Use:   "addr [private-key]",
	Short: "Get the address from a private key",
	Args:  cobra.ExactArgs(1),
	RunE:  getAddressFromPrivateKeyHandler,
}

func getAddressFromPrivateKeyHandler(cmd *cobra.Command, args []string) error {
	privKey, err := crypto.HexToECDSA(cmdhex.StripHexPrefix(args[0]))
	if err != nil {
		return err
	}
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	cmd.Println(addr.Hex())
	return nil
}
