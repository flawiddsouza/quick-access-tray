package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Command struct {
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
	Open    string `yaml:"open"`
	Group   string `yaml:"group"`
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

func splitString(s string) []string {
	stringToReturn := []string{}
	inQuotes := false
	currentString := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '"' {
			inQuotes = !inQuotes
			currentString += string(s[i])
		} else if s[i] == ' ' && !inQuotes {
			if currentString != "" {
				stringToReturn = append(stringToReturn, currentString)
				currentString = ""
			}
		} else {
			currentString += string(s[i])
		}
	}

	if currentString != "" {
		stringToReturn = append(stringToReturn, currentString)
	}

	return stringToReturn
}

func openFile(filename string) {
	println("Opening:", filename)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filename)
	case "windows":
		cmd = exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", filename)
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
