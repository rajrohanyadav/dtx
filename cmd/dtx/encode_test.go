package cmd

import (
	"bytes"
	"testing"
)

func TestEncodeCmd_Execute(t *testing.T) {
	cmd := newEncodeCmd()
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{"-t", "b64", "-s", "test"})

	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "dGVzdA==\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
