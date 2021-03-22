package controllers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/blockchain"
	. "server/db"
	"server/models"
	"server/utils"
)

type WalletResponse struct {
	UserID            uint    `json:"user_id"`
	PrivateKey        string  `json:"private_key"`
	PublicKey         string  `json:"public_key"`
	BlockchainAddress string  `json:"blockchain_address"`
	Amount            float32 `json:"amount"`
}

type TransactionRequest struct {
	SenderPrivateKey           string  `json:"sender_private_key"`
	SenderPublicKey            string  `json:"sender_public_key"`
	SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
	Value                      float32 `json:"value"`
}

func GetWallet(c *gin.Context) {
	uid := c.Param("user_id")
	if uid == "" {
		c.JSON(http.StatusBadRequest, "Please specify an user id.")
		return
	}

	w := models.Wallet{}
	err := DB.Where("user_id = ?", uid).Take(&w).Error
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found.")
		return
	}

	res := WalletResponse{
		UserID:            w.UserID,
		PrivateKey:        w.PrivateKey,
		PublicKey:         w.PublicKey,
		BlockchainAddress: w.BlockchainAddress,
		Amount:            w.Amount,
	}

	c.JSON(http.StatusOK, res)
}

func CreateTransaction(c *gin.Context) {
	tr := TransactionRequest{}

	err := c.BindJSON(&tr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	publicKey := utils.PublicKeyFromString(tr.SenderPublicKey)
	privateKey := utils.PrivateKeyFromString(tr.SenderPrivateKey, publicKey)

	// generate signiture
	m, _ := json.Marshal(tr)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, h[:])
	signature := &utils.Signature{r, s}

	// FIXME: Send request to blockchain node in p2p network
	/*
	   signatureStr := signature.String()
	   btr := &blockchain.TransactionRequest{
	       SenderBlockchainAddress: &tr.SenderBlockchainAddress,
	       RecipientBlockchainAddress: &tr.RecipientBlockchainAddress,
	       SenderPublicKey: &tr.SenderPublicKey,
	       Value: &tr.Value,
	       Signature: &signatureStr,
	   }
	*/
	bc := blockchain.GetBlockchain()
	created := bc.CreateTransaction(
		tr.SenderBlockchainAddress,
		tr.RecipientBlockchainAddress,
		tr.Value,
		publicKey,
		signature,
	)

	if !created {
		c.JSON(http.StatusBadRequest, "Faild to create transaction.")
		return
	}

	c.JSON(http.StatusCreated, "success!")
}
