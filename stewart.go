package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	// check if Terraform is installed
	var terraformExecutable string
	if runtime.GOOS == "windows" {
		terraformExecutable = "terraform.exe"
	} else {
		terraformExecutable = "terraform"
	}

	_, err := exec.LookPath(terraformExecutable)
	if err != nil {
		fail("terraform binary is missing")
	}

	// check if we received all required arguments
	if len(os.Args) < 3 {
		usage()
	}

	// check if required environment variable is set
	_, exists := os.LookupEnv("ARM_ACCESS_KEY")
	if !exists {
		fail("You need to set the environment variable ARM_ACCESS_KEY")
	}
}

func main() {
	subCommand := os.Args[1]
	environment := os.Args[2]

	cmd := fmt.Sprintf("terraform %s -var-file environments/%s.tfvars", subCommand, environment)

	if subCommand == "apply" {
		cmd += " -auto-approve"
	}

	initBackend()
	switchWorkspace(environment)

	info(fmt.Sprintf("executing: %s", cmd))
	streamCmdOutput(cmd)
}

func initBackend() {
	info("initializing backend")
	streamCmdOutput("terraform init")
}

func switchWorkspace(workspace string) {
	info(fmt.Sprintf("switching to workspace: %s", workspace))
	streamCmdOutput(fmt.Sprintf("terraform workspace select %s ", workspace))
}

func usage() {
	fmt.Println("usage: stewart $terraformSubcommand $environment")
	os.Exit(1)
}
