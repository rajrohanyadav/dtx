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
	"fmt"

	"github.com/rajrohanyadav/dtx/cmd/utils"
	"github.com/spf13/cobra"
)

func newEncodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encode",
		Short: "Encode [b64|jwt]",
		Long:  `Encode [b64|jwt]`,
		RunE: func(cmd *cobra.Command, args []string) error {
			op, _ := cmd.Flags().GetString("type")
			switch op {
			case "b64":
				{
					str, _ := cmd.Flags().GetString("str")
					res, err := utils.EncodeStringToB64(str)
					if err != nil {
						fmt.Println("Error converting to base 64")
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), res)
				}
			default:
				if err := cmd.Help(); err != nil {
					return err
				}
			}
			return nil
		},
	}
	addTypeFlag(cmd)
	addStringFlag(cmd)
	return cmd
}
