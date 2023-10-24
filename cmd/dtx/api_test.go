package cmd

import (
	"bytes"
	"testing"
)

func TestApiCmd_Execute(t *testing.T) {
	cmd := newAPICmd()
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
