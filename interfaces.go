package main

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type SignInResponse struct {
	DeepLinkURL  string `json:"deepLinkUrl"`
	PollingToken string `json:"pollingToken"`
	PublicKey    string `json:"publicKey"`
	PrivateKey   string `json:"privateKey"`
}

type WarpcastResponse struct {
	Result struct {
		SignedKeyRequest struct {
			Token      string `json:"token"`
			DeeplinkURL string `json:"deeplinkUrl"`
		} `json:"signedKeyRequest"`
	} `json:"result"`
}

var Domain = apitypes.TypedDataDomain{
	Name:              "Farcaster SignedKeyRequestValidator",
	Version:           "1",
	ChainId:           math.NewHexOrDecimal256(10),
	VerifyingContract: "0x00000000fc700472606ed4fa22623acf62c60553",
}

var Types = apitypes.Types{
	"EIP712Domain": []apitypes.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	},
	"SignedKeyRequest": []apitypes.Type{
		{Name: "requestFid", Type: "uint256"},
		{Name: "key", Type: "bytes"},
		{Name: "deadline", Type: "uint256"},
	},
}

type SignedKeyRequest struct {
	Token        string `json:"token"`
	DeeplinkUrl  string `json:"deeplinkUrl"`
	Key          string `json:"key"`
	RequestFid   int    `json:"requestFid"`
	State        string `json:"state"`
	IsSponsored  bool   `json:"isSponsored"`
	UserFid      int    `json:"userFid"`
}

type Result struct {
	SignedKeyRequest SignedKeyRequest `json:"signedKeyRequest"`
}

type PollResponse struct {
	Result Result `json:"result"`
}
