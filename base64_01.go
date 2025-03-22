package main

import (
	"encoding/base64"
	"fmt"
)

func DecodeKey(publickey string) string {
	publicKey, _ := base64.StdEncoding.DecodeString(publickey)

	fmt.Println("FromBase64String1: ", publicKey[1])
	fmt.Println("FromBase64String2: ", publicKey[2])
	fmt.Println("FromBase64String3: ", publicKey[3])
	en := base64.StdEncoding.EncodeToString([]byte(publicKey))
	fmt.Println("ToBase64String: ", en)

	return string(publicKey)
}
