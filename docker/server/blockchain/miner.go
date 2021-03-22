package blockchain

import (
	. "server/db"
	"server/models"
)

var cache map[string]*Blockchain = make(map[string]*Blockchain)

func GetBlockchain() *Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := models.NewWallet()
		DB.Save(&minersWallet)
		bc = NewBlockchain(minersWallet.BlockchainAddress)
		cache["blockchain"] = bc
	}
	return bc
}
