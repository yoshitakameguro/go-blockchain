package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "server/db"
	"server/test"
	"testing"
)

func init() {
	test.InitTest()
	test.R.GET("/wallet/:user_id", GetWallet)
	test.R.POST("/transaction", CreateTransaction)
}

func TestGetComment(t *testing.T) {
	user := FakeUserWithWallet()
	userID := fmt.Sprint(user.ID)
	endpoint := "/wallet/" + userID

	res, err := test.Request("GET", endpoint, nil)
	if err != nil {
		t.Errorf("Request failed: %v", err)
	}

	resMap := make(map[string]interface{})

	err = json.Unmarshal([]byte(res.Body.String()), &resMap)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, user.Wallet.UserID, uint(resMap["user_id"].(float64)))
	assert.Equal(t, user.Wallet.PrivateKey, resMap["private_key"])
	assert.Equal(t, user.Wallet.PublicKey, resMap["public_key"])
	assert.Equal(t, user.Wallet.BlockchainAddress, resMap["blockchain_address"])
	assert.Equal(t, user.Wallet.Amount, float32(resMap["amount"].(float64)))

	test.ClearDB()
}

func TestRequestTransaction(t *testing.T) {
	sender := FakeUserWithWallet()
	recipient := FakeUserWithWallet()
	endpoint := "/transaction"
	data, _ := json.Marshal(map[string]interface{}{
		"sender_private_key":           sender.Wallet.PrivateKey,
		"sender_public_key":            sender.Wallet.PublicKey,
		"sender_blockchain_address":    sender.Wallet.BlockchainAddress,
		"recipient_blockchain_address": recipient.Wallet.BlockchainAddress,
		"value":                        2.0,
	})

	res, err := test.Request("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Request failed: %v", err)
	}

	assert.Equal(t, 201, res.Code)

	test.ClearDB()
}
