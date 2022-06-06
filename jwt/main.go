package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	// Private & Public Key
	privateKeyPath := "app.rsa"
	publicKeyPath := "app.rsa.pub"

	privateKeyByte, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		panic(err)
	}

	publicKeyByte, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		panic(err)
	}

	// Init Claims
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Second)),
		},
		Username: "username",
		Email:    "test.test@test.test",
	}

	// Encode JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	string_token, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("-- Encode JWT Token --")
	fmt.Println(string_token)
	fmt.Println()

	// Decode JWT Token
	result_token, _ := jwt.ParseWithClaims(string_token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	// Validate JWT Token
	if !result_token.Valid {
		if result_token.Claims != nil {
			// Token is expired
			fmt.Println("Result: Token is expired")
			return
		} else {
			// Token is not valid
			fmt.Println("Result: Token is not valid")
			return
		}
	}

	// Get JWT Token Claim
	fmt.Println("-- Decode JWT Token --")
	fmt.Printf("Email: %s\n", result_token.Claims.(*MyClaims).Email)
}
