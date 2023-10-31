package cmd

import (
	"bytes"
	"testing"
)

func TestGenerateCmd_Execute(t *testing.T) {
	cmd := newGenerateCmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Generate the specified resource.\n\nUsage:\n  generate [flags]\n\nFlags:\n  -n, --count int     number of outputs required (default 1)\n  -h, --help          help for generate\n  -t, --type string   type of input [b64|jwt]\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
