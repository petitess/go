package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "sdpofj21323roiwssdvsvvsdvsvsdvsvsdvsdvd"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 4).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VarifiyToken(token string) (int64, error) {
	// Remove escaped double quotes from the token string
	token = strings.ReplaceAll(token, "\"", "")

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, errors.New("invalid token format")
	}

	// Decode the header
	_, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return 0, fmt.Errorf("could not decode header: %w", err)
	}

	// Decode the payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, fmt.Errorf("could not decode payload: %w", err)
	}

	// Parse the payload as a JSON map
	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return 0, fmt.Errorf("could not parse payload: %w", err)
	}
	// Extract the userId from the payload
	userId, ok := payload["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid userId in token")
	}

	return int64(userId), nil
}

// func VarifiyToken(token string) (int64, error) {
// 	token = strings.ReplaceAll(token, "\"", "")
// 	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 		if !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		return 0, errors.New("could not parse token. " + err.Error())
// 	}
// 	tokenIsValid := parsedToken.Valid
// 	if tokenIsValid {
// 		return 0, errors.New("invalid token")
// 	}
// 	claims, ok := parsedToken.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return 0, errors.New("invalid token claims")

// 	}

// 	// email := claims["email"].(string)
// 	userId := claims["userId"].(int64)
// 	return userId, nil
// }
