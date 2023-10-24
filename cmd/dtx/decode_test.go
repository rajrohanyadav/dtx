package cmd

import (
	"bytes"
	"testing"
)

func TestDecodeCmd_Execute(t *testing.T) {
	cmd := newDecodeCmd()
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{"-t", "b64", "-s", "dGVzdA=="})

	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "test\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
