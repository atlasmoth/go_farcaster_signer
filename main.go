package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func createSigner(c *gin.Context) {
	signInData, err := SignInWithWarpcast()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to sign in user: %v", err)})
		return
	}

	response := SignInResponse{
		DeepLinkURL:  signInData["signerApprovalUrl"].(string),
		PollingToken: signInData["token"].(string),
		PublicKey:    signInData["publicKey"].(string),
		PrivateKey:   signInData["privateKey"].(string),
	}

	c.JSON(http.StatusOK, response)
}

func getSignerStatus(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	result, err := GetSignerFromWarpcast(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get signer: %v", err)})
		return
	}

	

	c.JSON(http.StatusOK, result)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.POST("/signer", createSigner)
	r.GET("/signer-status",getSignerStatus)

	log.Println("Server is running on :8080")
	r.Run(":8080")
}