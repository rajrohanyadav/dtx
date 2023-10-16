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

	"github.com/rajrohanyadav/dtx/cmd/utils"
	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode [b64|jwt]",
	Long:  `Encode [b64|jwt]`,
	Run: func(cmd *cobra.Command, args []string) {
		op, _ := cmd.Flags().GetString("type")
		if op == "b64" {
			str, _ := cmd.Flags().GetString("str")
			res, err := EncodeStringToB64(str)
			if err != nil {
				fmt.Println("Error converting to base 64")
			}
			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	utils.AddTypeFlag(encodeCmd)
	utils.AddStringFlag(encodeCmd)
}

func EncodeStringToB64(str string) (string, error) {
	return b64.StdEncoding.EncodeToString([]byte(str)), nil
}
