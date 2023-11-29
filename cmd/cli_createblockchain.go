package cmd

import (
	"Blockchainlab/pkg"
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !pkg.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := pkg.CreateBlockchain(address, nodeID)
	defer bc.Db.Close()

	UTXOSet := pkg.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
