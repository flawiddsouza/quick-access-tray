package main

import (
	"os/exec"
)

func runCommand(command string) {
	println("Running command:", command)
	split_command := splitString(command)
	cmd := exec.Command(split_command[0], split_command[1:]...)
	err := cmd.Run()
	if err != nil {
		println("Failed to run command:", err)
	} else {
		println("Command ran successfully")
	}
}
