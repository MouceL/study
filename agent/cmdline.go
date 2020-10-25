package agent

import (
	"os"
	"os/exec"
	"strings"
)

func ExecuteCommand(args string) bool{
	cmdArray := strings.Split(args," ")
	cmd := exec.Command(cmdArray[0],cmdArray[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
	}
	// wait
	return true
}
