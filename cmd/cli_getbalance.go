package cmd

import (
	"Blockchainlab/pkg"
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !pkg.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := pkg.NewBlockchain(nodeID)
	UTXOSet := pkg.UTXOSet{bc}
	defer bc.Db.Close()

	balance := 0
	pubKeyHash := pkg.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
