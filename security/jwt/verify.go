package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// VerifyToken verifies the authenticity of a JWT token using a public key.
// It takes a token string as input and returns a parsed JWT token along with any errors encountered during the process.
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Read the public key file specified in the SECURITY_PUBLIC_KEY environment variable.
	publicKeyBytes, err := ioutil.ReadFile(os.Getenv("SECURITY_PUBLIC_KEY"))
	if err != nil {
		return nil, err
	}

	// Decode the PEM-encoded public key.
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("decode key error")
	}

	// Parse the public key from the decoded bytes.
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("decode key error")
	}

	// Parse and verify the JWT token using the public key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is RSA.
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Token method is not RSA")
		}
		// Return the public key for verification.
		return publicKey.(*rsa.PublicKey), nil
	})

	// Return the parsed token and any error encountered during verification.
	return token, err
}