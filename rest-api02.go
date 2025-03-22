package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/nacl/box"
)

var generateKey = box.GenerateKey

const (
	baseUrl     = "https://api.github.com/repos/petitess/repo-dev"
	owner       = "petitess"                                 // Replace with the repository owner
	repo        = "repo-dev"                                 // Replace with the repository name
	secretName  = "MY_SECRET"                                // Replace with the secret name
	secretValue = "1234abcdRUMBA12"                          // Replace with the secret value
	token       = "ghp_xxxx" // Replace with your GitHub token
	keySize     = 32
	nonceSize   = 24
)

func main() {
	public_key, public_key_id := GetGhPublicKey("/actions/secrets/public-key")
	fmt.Println("public-key: " + public_key[1:10])
	fmt.Println("public-key_id: " + public_key_id[1:10])

	sealed, _ := Encrypt(public_key, secretValue)
	createOrUpdateSecret(sealed, public_key_id)

}

func GetGhPublicKey(url string) (string, string) {
	type PublicKey struct {
		Key_id string `json:"key_id"`
		Key    string `json:"key"`
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	resp, _ := client.Do(req)
	// resp, _ := http.DefaultClient.Do(req)
	// fmt.Printf("client: status code: %d\n", resp.StatusCode)
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var publicKey PublicKey
	json.Unmarshal(responseData, &publicKey)
	// fmt.Printf("%s\n", responseData)
	// fmt.Println(publicKey.Key)
	return publicKey.Key, publicKey.Key_id
}

func createOrUpdateSecret(encryptedValue string, key_id string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/secrets/%s", owner, repo, secretName)
	// encryptedValue = "OC1A8RhXwICuPdomg8358h1hZy8Wq2MJj8mh1g062CQ2FIQFA2Suxr1sNgQw4FdyY2n2/sqSX+k="
	// Create the payload
	payload := map[string]interface{}{
		"encrypted_value": encryptedValue, // This should be the encrypted value of your secret
		"key_id":          key_id,         // Replace with the key ID used for encryption
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create a new request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.raw+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	// req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err.Error())
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	fmt.Println("Secret created or updated successfully! " + string(resp.Status))
}

func createVariable(encryptedValue string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/variables", owner, repo)
	//url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/variables/%s", owner, repo, secretName)

	// Create the payload
	payload := map[string]interface{}{
		"value": encryptedValue, // This should be the encrypted value of your secret
		"name":  secretName,     // Replace with the key ID used for encryption
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.raw+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	// if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("Error: received status code %d", resp.StatusCode)
	// }

	fmt.Println("Variable created successfully! " + string(resp.StatusCode))
}

func updateVariable(encryptedValue string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/variables/%s", owner, repo, secretName)

	// Create the payload
	payload := map[string]interface{}{
		"value": encryptedValue, // This should be the encrypted value of your secret
		"name":  secretName,     // Replace with the key ID used for encryption
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create a new request
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.raw+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	// if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("Error: received status code %d", resp.StatusCode)
	// }

	fmt.Println("Variable updated successfully! " + string(resp.StatusCode))
}

// Encrypt encrypts a secret using the provided recipient public key.
func Encrypt(recipientPublicKey string, content string) (string, error) {
	// decode the provided public key from base64
	recipientKey := new([keySize]byte)
	b, err := base64.StdEncoding.DecodeString(recipientPublicKey)
	if err != nil {
		return "", err
	} else if size := len(b); size != keySize {
		return "", fmt.Errorf("recipient public key has invalid length (%d bytes)", size)
	}

	copy(recipientKey[:], b)

	// create an ephemeral key pair
	pubKey, privKey, err := generateKey(rand.Reader)
	if err != nil {
		return "", err
	}

	// create the nonce by hashing together the two public keys
	nonce := new([nonceSize]byte)
	nonceHash, err := blake2b.New(nonceSize, nil)
	if err != nil {
		return "", err
	}

	if _, err := nonceHash.Write(pubKey[:]); err != nil {
		return "", err
	}

	if _, err := nonceHash.Write(recipientKey[:]); err != nil {
		return "", err
	}

	copy(nonce[:], nonceHash.Sum(nil))

	// begin the output with the ephemeral public key and append the encrypted content
	out := box.Seal(pubKey[:], []byte(content), nonce, recipientKey, privKey)

	// base64-encode the final output
	return base64.StdEncoding.EncodeToString(out), nil
}
