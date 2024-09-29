package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

var newAcc = &cobra.Command{
	Use:   "new-acc",
	Short: "Create a random account",
	Args:  cobra.ExactArgs(0),
	RunE:  newAccHandler,
}

var getAddressFromPrivateKey = &cobra.Command{
	Use:   "addr [private-key]",
	Short: "Get the address from a private key",
	Args:  cobra.ExactArgs(1),
	RunE:  getAddressFromPrivateKeyHandler,
}

func init() {
	var RPCAddr string
	getBalanceCmd.PersistentFlags().StringVarP(&RPCAddr, "rpc-addr", "r", DefaultRPCAddr, "The RPC address of the Ethereum node")
	ethCmd.AddCommand(getBalanceCmd)
	ethCmd.AddCommand(newAcc)
	ethCmd.AddCommand(getAddressFromPrivateKey)
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
	fmt.Printf("%s\n", marshaled)
	return nil
}

func getAddressFromPrivateKeyHandler(cmd *cobra.Command, args []string) error {
	privKey, err := crypto.HexToECDSA(args[0])
	if err != nil {
		return err
	}
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	fmt.Println(addr.Hex())
	return nil
}
