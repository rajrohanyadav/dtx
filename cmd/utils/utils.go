package utils

import (
	b64 "encoding/base64"

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
