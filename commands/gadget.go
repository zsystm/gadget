package commands

import (
	"github.com/spf13/cobra"
)

var (
	GadgetCmd = &cobra.Command{
		Use:   "gadget",
		Short: "gadget is a swiss army knife for cosmos sdk based chains",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func init() {
	addCommands()
}

func Execute() {
	if err := GadgetCmd.Execute(); err != nil {
		panic(err)
	}
}

func addCommands() {
	GadgetCmd.AddCommand(EthCmd)
	GadgetCmd.AddCommand(BechToEthCmd)
	GadgetCmd.AddCommand(OtherPrefixCmd)
	GadgetCmd.AddCommand(EthToBechCmd)
	GadgetCmd.AddCommand(Base64ToHexCmd)
	GadgetCmd.AddCommand(HexToBase64Cmd)
	GadgetCmd.AddCommand(Secp256k1Cmd)
}
