package utils

import "github.com/spf13/cobra"

var typeFlag string
var stringFlag string

func AddTypeFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&typeFlag, "type", "t", "", "type of input [b64|jwt]")
}

func AddStringFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&stringFlag, "str", "s", "", "string input")
}
