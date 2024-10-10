package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func SignInWithWarpcast() (map[string]interface{}, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %v", err)
	}
	keypairString := map[string]string{
		"publicKey":  "0x" + hex.EncodeToString(publicKey),
		"privateKey": "0x" + hex.EncodeToString(privateKey),
	}

	mnemonic := os.Getenv("APP_MNEMONIC")
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	account, err := wallet.Derive(accounts.DefaultBaseDerivationPath, true)

	appFid := os.Getenv("APP_FID")

	deadline := time.Now().Unix() + 86400 // signature is valid for 1 day

	deadlineToBigInt := new(big.Int).SetInt64(deadline)
	
	typedData := apitypes.TypedData{
		Types:Types,
		PrimaryType: "SignedKeyRequest",
		Domain:      Domain,
		Message: map[string]interface{}{
			"requestFid": appFid,
			"key":        keypairString["publicKey"],
			"deadline":  deadlineToBigInt,
		},
	}

	
	
	signature, err := SignEIP712TypedData(wallet, account, typedData)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	
	
	authData := map[string]interface{}{
		"signature":     "0x" + hex.EncodeToString(signature),
		"requestFid":    appFid,
		"deadline":      deadlineToBigInt,
		"requestSigner":account.Address.Hex(),
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"key":        keypairString["publicKey"],
		"signature":  authData["signature"],
		"requestFid": appFid,
		"deadline":   deadlineToBigInt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	resp, err := http.Post("https://api.warpcast.com/v2/signed-key-requests", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var warpcastResp WarpcastResponse
	if err := json.Unmarshal(body, &warpcastResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	user := map[string]interface{}{
		"publicKey":         keypairString["publicKey"],
		"deadline":          deadline,
		"token":             warpcastResp.Result.SignedKeyRequest.Token,
		"signerApprovalUrl": warpcastResp.Result.SignedKeyRequest.DeeplinkURL,
		"privateKey":        keypairString["privateKey"],
		"status":            "pending_approval",
	}

	for k, v := range authData {
		user[k] = v
	}

	return user, nil
}