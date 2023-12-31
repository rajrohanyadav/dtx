package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestApiCmd_Help(t *testing.T) {
	cmd := newAPICmd()
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

func TestApiCmd_Execute(t *testing.T) {
	cmd := newAPICmd()
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// TODO: use mocks instead
	cmd.SetArgs([]string{"-t", "get", "-s", "https://jsonplaceholder.typicode.com/todos/1"})

	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Output:\n{\n  \"completed\": false,\n  \"id\": 1,\n  \"title\": \"delectus aut autem\",\n  \"userId\": 1\n}\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, not contained in: %q", cmd.Long, stdout.String())
	}
}
