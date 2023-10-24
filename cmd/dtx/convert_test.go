package cmd

import (
	"bytes"
	"testing"
)

func TestCoverageCmd_Execute(t *testing.T) {
	cmd := newConvertCmd()
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
