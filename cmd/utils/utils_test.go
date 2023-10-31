package utils

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockHTTPClient struct{}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return httptest.NewRecorder().Result(), nil
}

func (m *mockHTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	return httptest.NewRecorder().Result(), nil
}

type mockHTTPClientWithError struct{}

func (m *mockHTTPClientWithError) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("custom error from mock client")
}

func (m *mockHTTPClientWithError) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("custom error from mock client")
}

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
	uuids, err := GenerateUUID(5)

	if err != nil {
		t.Errorf("Error generating UUIDs: %v", err)
	}

	if len(uuids) != 5 {
		t.Errorf("Expected 5 UUIDs, but got %d", len(uuids))
	}

	for _, u := range uuids {
		_, err := uuid.Parse(u)
		if err != nil {
			t.Errorf("Invalid UUID: %v", err)
		}
	}
}

func TestMakeGetRequest(t *testing.T) {
	mockClient := &http.Client{
		Transport: &mockHTTPClient{},
	}

	http.DefaultClient = mockClient

	uri := "https://example.com"
	res, err := MakeGetRequest(uri)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	invalidURI := "invalid-url"
	http.DefaultClient = &http.Client{
		Transport: &mockHTTPClientWithError{},
	}
	_, err = MakeGetRequest(invalidURI)
	if err == nil {
		t.Error("Expected an error from the mock HTTP client")
	}
}
