package cmd

import (
	"Blockchainlab/pkg"
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !pkg.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !pkg.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := pkg.NewBlockchain(nodeID)
	UTXOSet := pkg.UTXOSet{bc}
	defer bc.Db.Close()

	wallets, err := pkg.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := pkg.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := pkg.NewCoinbaseTX(from, "")
		txs := []*pkg.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		//pkg.sendTx(pkg.knownNodes[0], tx)
		pkg.SendTx(pkg.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
