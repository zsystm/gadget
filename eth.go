package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
)

const (
	DefaultRPCAddr = "http://localhost:8545"
)

var ethCmd = &cobra.Command{
	Use:   "eth",
	Short: "Ethereum related commands",
	Long:  "All commands related to Ethereum",
}

var getBalanceCmd = &cobra.Command{
	Use:   "get-balance [address]",
	Short: "Get the balance of an account",
	Args:  cobra.ExactArgs(1),
	RunE:  getBalanceCmdHandler,
}

func init() {
	var RPCAddr string
	getBalanceCmd.PersistentFlags().StringVarP(&RPCAddr, "rpc-addr", "r", DefaultRPCAddr, "The RPC address of the Ethereum node")
	ethCmd.AddCommand(getBalanceCmd)
}

func getBalanceCmdHandler(cmd *cobra.Command, args []string) error {
	rpcAddr, _ := cmd.Flags().GetString("rpc-addr")

	addrStr := common.HexToAddress(args[0])
	addr := addrStr.Hex()
	if !common.IsHexAddress(addr) {
		return fmt.Errorf("invalid address: %s", addr)
	}

	ethCli, err := ethclient.Dial(rpcAddr)
	if err != nil {
		return err
	}
	defer ethCli.Close()

	balance, err := ethCli.BalanceAt(context.Background(), common.HexToAddress(args[0]), nil)
	if err != nil {
		return err
	}
	fmt.Println(balance.String())
	return nil
}
