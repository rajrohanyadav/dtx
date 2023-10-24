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

func newGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate the specified resource",
		Long:  `Generate the specified resource.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			typ, _ := cmd.Flags().GetString("type")
			if typ == "uuid" {
				num, _ := cmd.Flags().GetInt("count")
				res, err := utils.GenerateUUID(num)
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, s := range res {
					fmt.Println(s)
				}
			}
			return nil
		},
	}
	utils.AddTypeFlag(cmd)
	utils.AddNumberFlag(cmd)
	return cmd
}
