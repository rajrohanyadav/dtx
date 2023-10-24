package cmd

import (
	"bytes"
	"testing"

	// "github.com/stretchr/testify/assert"
)

func TestRootCmd_Execute(t *testing.T) {
	cmd := newRootCommand()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	// cmd.SetArgs([]string{""})
	cmd.Execute()

	// expected := "Dev 10X"

	// assert.Equal(t, actual.String(), expected, "actual is not expected")
}
