// Package jwt provides functionality for handling JSON Web Tokens (JWT).
package jwt

import (
	"net/http"
	"os"
	"sayamphoo/microservice/models/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// BuildJwt generates a JWT token with the provided subject (sub) using RSA-256 signing method.
// It reads the private key from the file specified in the SECURITY_PRIVATE_KEY environment variable.
// The generated token includes claims for issuer (iss), subject (sub), and expiration time (exp).
// The token is signed using the RSA private key and returned as a string.
func BuildJwt(sub *string) (*string, error) {
	// Retrieve the path to the private key file from the environment variable.
	privateKeyFile := os.Getenv("SECURITY_PRIVATE_KEY")

	// Read the private key file content.
	privateKeyBytes, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	// Create a new JWT token with specified claims and signing method.
	tos := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iss": "my-auth-server",                           // Issuer claim
			"sub": sub,                                        // Subject claim
			"exp": time.Now().Add(45 * 24 * time.Hour).Unix(), // Expiration time claim (45 days)
		})

	// Parse the RSA private key from the PEM-encoded key bytes.
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		// Handle error if unable to parse the private key.
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "failed to read private key file",
		})
	}

	// Sign the token with the private key and get the signed token as a string.
	signedToken, err := tos.SignedString(key)
	if err != nil {
		// Handle error if unable to sign the token.
		panic(domain.UtilityModel{
			Code:    http.StatusInternalServerError,
			Message: "failed to parse private key",
		})
	}

	// Return the signed JWT token.
	return &signedToken, nil
}
