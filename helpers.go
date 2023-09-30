package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

func getIcon(iconName string) []byte {
	iconPath, _ := filepath.Abs(iconName)
	trayIcon, err := os.Open(iconPath)
	if err != nil {
		panic(err)
	}
	defer trayIcon.Close()

	iconStat, err := trayIcon.Stat()
	if err != nil {
		panic(err)
	}
	iconBytes := make([]byte, iconStat.Size())
	_, err = trayIcon.Read(iconBytes)
	if err != nil {
		panic(err)
	}
	return iconBytes
}
