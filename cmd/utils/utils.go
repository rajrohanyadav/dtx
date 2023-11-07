/*
Copyright © 2023 Rohan Yadav rajrohanyadav@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
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

func WriteOutput(out io.Writer, resp string, err error) {
	warn := color.New(color.Bold, color.FgRed).FprintlnFunc()
	success := color.New(color.Bold, color.FgGreen).FprintlnFunc()
	if err != nil {
		warn(out, "Error:")
		fmt.Fprintf(out, "something went wrong: %s\n", err)
	} else {
		success(out, "Output:")
		fmt.Fprintf(out, "%s\n", resp)
	}
}
