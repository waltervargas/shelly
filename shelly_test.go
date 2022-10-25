package shelly_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/shelly"
)

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	input := "/bin/ls -l main.go"
	want := []string{"/bin/ls", "-l", "main.go"}
	cmd, err := shelly.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromStringNoArgs(t *testing.T) {
	t.Parallel()
	input := "/bin/ls"
	want := []string{"/bin/ls"}
	cmd, err := shelly.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shelly.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}
