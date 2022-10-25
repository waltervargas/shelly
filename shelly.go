package shelly

import (
	"fmt"
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
