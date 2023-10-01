package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"gopkg.in/yaml.v3"
)

type Command struct {
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
}

func parseConfigYAML(configFilePath string) ([]Command, error) {
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config []Command
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func runCommand(command string) {
	println("Running command:", command)
	split_command := strings.Split(command, " ")
	cmd := exec.Command(split_command[0], split_command[1:]...)
	err := cmd.Run()
	if err != nil {
		println("Failed to run command:", err)
	} else {
		println("Command ran successfully")
	}
}

func openFile(filename string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filename)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filename)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // so that a cmd window doesn't pop up
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
