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
package cmd

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		op, _ := cmd.Flags().GetString("type")
		if op == "b64" {
			b64Str, _ := cmd.Flags().GetString("str")
			res, err := DecodeB64ToString(b64Str)
			if err != nil {
				fmt.Println("Error decoding from base 64")
			}
			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	encodeCmd.Flags().StringP("type", "t", "b64", "What do you want to encode to")
	encodeCmd.Flags().StringP("str", "s", "", "string to decode")
}

func DecodeB64ToString(b64Str string) (string, error) {
	res, err := b64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
