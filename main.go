package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gadget",
	Short: "A CLI tool for various tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand. See 'gadget --help' for more details.")
	},
}

func main() {
	initializeCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initializeCommands() {
	rootCmd.AddCommand(ethCmd)
	rootCmd.AddCommand(bechToEthCmd)
	rootCmd.AddCommand(ethToBechCmd)
	rootCmd.AddCommand(base64ToHexCmd)
	rootCmd.AddCommand(hexToBase64Cmd)
}
