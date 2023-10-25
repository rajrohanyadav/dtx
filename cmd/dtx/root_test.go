package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCmd_Execute(t *testing.T) {
	cmd := newRootCommand()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if strings.Contains(stdout.String(), cmd.Long) {
		t.Errorf("Expected output: %q, not contained in: %q", cmd.Long, stdout.String())
	}
}
