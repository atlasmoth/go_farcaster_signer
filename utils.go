package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)


func SignEIP712TypedData(wallet *hdwallet.Wallet, account accounts.Account, typedData apitypes.TypedData) ([]byte, error) {

	hash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return nil, err
	}

	signature, err := wallet.SignHash(account, hash)
	if err != nil {
		return nil, err
	}

	signature[64] += 27

	return signature, nil
}

func GetSignerFromWarpcast(token string) (*PollResponse, error) {
	url := fmt.Sprintf("https://api.warpcast.com/v2/signed-key-request?token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var result PollResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}