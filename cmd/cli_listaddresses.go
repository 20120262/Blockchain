package cmd

import (
	"Blockchainlab/pkg"
	"fmt"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := pkg.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
