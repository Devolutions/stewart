package main

import (
	"os"
	"os/exec"
	"strings"
)

func streamCmdOutput(command string) {
	splitCmd := strings.Split(command, " ")
	cmd := exec.Command(splitCmd[0], splitCmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fail(err.Error())
	}
}
