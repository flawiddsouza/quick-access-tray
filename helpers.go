package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func parseProjectsYAML() ([]Project, error) {
	yamlFile, err := os.ReadFile("projects.yml")
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = yaml.Unmarshal(yamlFile, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func openProject(command string) {
	split_command := strings.Split(command, " ")
	cmd := exec.Command(split_command[0], split_command[1:]...)
	err := cmd.Run()
	if err != nil {
		println("Failed to open project:", err)
	} else {
		println("Project opened successfully")
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
