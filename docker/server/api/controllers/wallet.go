package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "server/models"
    . "server/db"
)

type WalletResponse struct {
    UserID        uint      `json:"user_id"`
    PrivateKey    string `json:"private_key"`
    PublicKey     string  `json:"public_key"`
    BlockchainAddress  string    `json:"blockchain_address"`
    Amount float32 `json:"amount"`
}

func GetWallet(c *gin.Context) {
    userID := c.Param("user_id")
    if userID == "" {
        c.JSON(http.StatusBadRequest, "Please specify an user id.")
        return
    }

    wallet := models.Wallet{}
    err := DB.Where("user_id = ?", userID).Take(&wallet).Error
    if err != nil {
        c.JSON(http.StatusNotFound, "Not Found.")
        return
    }

    response := WalletResponse{
        UserID: wallet.UserID,
        PrivateKey: wallet.PrivateKey,
        PublicKey: wallet.PublicKey,
        BlockchainAddress: wallet.BlockchainAddress,
        Amount: wallet.Amount,
    }

    c.JSON(http.StatusOK, response)
}
