package commands

import (
	"github.com/spf13/cobra"

	"github.com/zsystm/gadget/commands/eth"
)

const (
	DefaultRPCAddr = "http://localhost:8545"
)

var EthCmd = &cobra.Command{
	Use:   "eth",
	Short: "eth is a tool for interacting with Ethereum",
}

func init() {
	var RPCAddr string
	eth.GetBalanceCmd.PersistentFlags().StringVarP(&RPCAddr, "rpc-addr", "r", DefaultRPCAddr, "The RPC address of the Ethereum node")
	EthCmd.AddCommand(eth.GetBalanceCmd)
	EthCmd.AddCommand(eth.NewAcc)
	EthCmd.AddCommand(eth.GetAddressFromPrivateKey)
}
