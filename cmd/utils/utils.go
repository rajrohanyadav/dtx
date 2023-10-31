package utils

import (
	b64 "encoding/base64"
	"net/http"

	"github.com/google/uuid"
)

func EncodeStringToB64(str string) (string, error) {
	return b64.StdEncoding.EncodeToString([]byte(str)), nil
}

func DecodeB64ToString(b64Str string) (string, error) {
	res, err := b64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func GenerateUUID(n int) ([]string, error) {
	var uuids []string
	for i := 0; i < n; i++ {
		id := uuid.New()
		uuids = append(uuids, id.String())
	}
	return uuids, nil
}

func MakeGetRequest(uri string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
