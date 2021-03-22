package controllers

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
    "server/test"
    . "server/db"
)

func init() {
    test.InitTest()
	test.R.GET("/wallet/:user_id", GetWallet)
}

func TestGetComment(t *testing.T) {
	user := FakeUserWithWallet()
	userID := fmt.Sprint(user.ID)
	endpoint := "/wallet/" + userID

    fmt.Println(endpoint)

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
