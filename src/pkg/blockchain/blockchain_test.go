package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFromAddress(t *testing.T) {
	// Define the structure for test cases
	type testCase struct {
		name          string
		generateKey   func() *ecdsa.PrivateKey
		expectPanic   bool
		expectedPanic string
	}

	// Initialize test cases
	testCases := []testCase{
		{
			name: "Valid ECDSA Private Key",
			generateKey: func() *ecdsa.PrivateKey {
				// Generate a new ECDSA private key using the P256 curve
				privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
				if err != nil {
					t.Fatalf("Failed to generate private key: %v", err)
				}
				return privateKey
			},
			expectPanic:   false,
			expectedPanic: "",
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			privateKey := tc.generateKey()
			expectedAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
			resultAddress := getFromAddress(privateKey)
			assert.Equal(t, expectedAddress, resultAddress, "The returned address should match the expected address")

		})
	}
}
