package cmd

import (
	"Blockchainlab/pkg"
	"fmt"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := pkg.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
