package shelly

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func CmdFromString(cmd string) (*exec.Cmd, error) {
	args := strings.Fields(cmd)
	if len(args) < 1 {
		return nil, fmt.Errorf("unable to create a command from empty string: %s", cmd)
	}
	return exec.Command(args[0], args[1:]...), nil
}

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	DryRun         bool
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		DryRun: false,
	}
}
