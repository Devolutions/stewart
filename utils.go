package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getCmdOutput(command string) []byte {
	splitCmd := strings.Split(command, " ")
	cmd := exec.Command(splitCmd[0], splitCmd[1:]...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		fail(fmt.Sprintf("command '%s' failed with: %s", command, err.Error()))
	}

	return output
}

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
