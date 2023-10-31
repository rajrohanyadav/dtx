package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeStringToB64(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "encode test",
			input:          "test",
			expectedOutput: "dGVzdA==",
			expectedError:  nil,
		},
		{
			name:           "encode user:password",
			input:          "user:password",
			expectedOutput: "dXNlcjpwYXNzd29yZA==",
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf(tt.name)
			res, err := EncodeStringToB64(tt.input)
			assert.Equal(t, tt.expectedOutput, res)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestDecodeB64ToString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "decode to test",
			input:          "dGVzdA==",
			expectedOutput: "test",
			expectedError:  nil,
		},
		{
			name:           "decode to user:password",
			input:          "dXNlcjpwYXNzd29yZA==",
			expectedOutput: "user:password",
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf(tt.name)
			res, err := DecodeB64ToString(tt.input)
			assert.Equal(t, tt.expectedOutput, res)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
		})
	}
}

func TestMakeGetRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()
	res, err := MakeGetRequest(mockServer.URL)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
