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
	"encoding/json"
	"fmt"
	"io"

	"github.com/rajrohanyadav/dtx/cmd/utils"
	"github.com/spf13/cobra"
)

func newAPICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api [flags] url",
		Short: "api [get|post|put|delete]",
		Long: `api [get|post|put|delete]
			Only GET is implemented, others are not implemented yet.
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			op, _ := cmd.Flags().GetString("type")
			switch op {
			case "get":
				{
					url, _ := cmd.Flags().GetString("str")
					resp, err := getAPI(url)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), resp)
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

func getAPI(url string) (string, error) {
	resp, err := utils.MakeGetRequest(url)
	if err != nil {
		return "", err
	}
	// TODO: check headers to see what kind of content we are getting
	// TODO: if json -> print or save json
	// TODO: if xml -> print or save xml
	// TODO: if file -> save file
	defer resp.Body.Close()
	responseJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data interface{}
	err = json.Unmarshal(responseJSON, &data)
	if err != nil {
		return "", err
	}
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}
