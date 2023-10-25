package main

import (
	"os/exec"
	"runtime"
	"syscall"
)

func runCommand(command string) {
	println("Running command:", command)
	split_command := splitString(command)
	cmd := exec.Command(split_command[0], split_command[1:]...)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // so that a cmd window doesn't pop up
	}
	err := cmd.Run()
	if err != nil {
		println("Failed to run command:", err)
	} else {
		println("Command ran successfully")
	}
}
