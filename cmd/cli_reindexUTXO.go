package cmd

import (
	"Blockchainlab/pkg"
	"fmt"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := pkg.NewBlockchain(nodeID)
	UTXOSet := pkg.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
