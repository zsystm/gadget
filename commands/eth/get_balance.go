package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var GetBalanceCmd = &cobra.Command{
	Use:   "get-balance [address]",
	Short: "Get the balance of an account",
	Args:  cobra.ExactArgs(1),
	RunE:  getBalanceCmdHandler,
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
	cmd.Println(balance.String())
	return nil
}
