package shelly

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprintf(s.Stdout, "> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nBe seeing you!")
			break
		}
		cmd, err := CmdFromString(line)
		if err != nil {
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s", output)
	}

}

func RunCLI() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
