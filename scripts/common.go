package scripts

import (
	"fmt"
	"os"
	"os/exec"
)

func handle(err error)  {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Pipe(where string, command string, args ...string)  {
	cmd := exec.Command(command, args...)
	cmd.Dir = where
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	handle(cmd.Run())
}
