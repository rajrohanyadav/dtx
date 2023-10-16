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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// encodeCmd.Flags().StringP("type", "t", "b64", "What do you want to encode to")
	// encodeCmd.Flags().StringP("str", "s", "", "string to encode")
	utils.AddTypeFlag(encodeCmd)
	utils.AddStringFlag(encodeCmd)
}

func EncodeStringToB64(str string) (string, error) {
	return b64.StdEncoding.EncodeToString([]byte(str)), nil
}
