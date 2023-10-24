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

	expectedOutput := ""
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
