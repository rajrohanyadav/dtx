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

	expectedOutput := "convert from t1 to t2. Not implemented yet\n\nUsage:\n  convert [flags]\n\nFlags:\n  -h, --help   help for convert\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
