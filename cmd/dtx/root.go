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
	"os"

	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dtx",
		Short: "Dev 10X",
		Long: `
		Dev 10X provides a suit of tools commonly needed by a develop.
		Easy to use in the terminal as standalone utility or piping outputs from other commands.
		Go to https://rajrohanyadav.github.io/dtx/ for detailed examples and usecases.
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	encodeCmd := newEncodeCmd()
	decodeCmd := newDecodeCmd()
	convertCmd := newConvertCmd()
	apiCmd := newAPICmd()
	generateCmd := newGenerateCmd()

	cmd.AddCommand(encodeCmd)
	cmd.AddCommand(decodeCmd)
	cmd.AddCommand(convertCmd)
	cmd.AddCommand(apiCmd)
	cmd.AddCommand(generateCmd)

	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := newRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
