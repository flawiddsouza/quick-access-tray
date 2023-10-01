package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

var configFilePath string

func createMenu() {
	var config []Command

	args := os.Args[1:]
	err := error(nil)

	if len(args) > 0 {
		if args[0] == "--config" {
			if len(args) > 1 {
				config, err = parseConfigYAML(args[1])
				if err != nil {
					println("Failed to parse config.yml:", err)
					systray.Quit()
				}
			} else {
				println("No config file specified.")
				systray.Quit()
			}
		}

		file_path, err := filepath.Abs(args[1])
		if err != nil {
			fmt.Println("Failed to get absolute path:", err)
			systray.Quit()
		}

		configFilePath = file_path
		fmt.Printf("Loaded config from %s\n", file_path)
	} else {
		file_path, err := filepath.Abs("config.yml")
		if err != nil {
			fmt.Println("Failed to get absolute path:", err)
			systray.Quit()
		}

		config, err = parseConfigYAML(file_path)
		if err != nil {
			println("Failed to parse config.yml:", err)
			systray.Quit()
		}

		configFilePath = file_path
		fmt.Printf("Loaded config from %s\n", file_path)
	}

	for _, command := range config {
		menuItem := systray.AddMenuItem(command.Label, command.Label)
		go func(command string) {
			for {
				select {
				case <-menuItem.ClickedCh:
					runCommand(command)
				}
			}
		}(command.Command)
	}

	systray.AddSeparator()

	mOpenConfig := systray.AddMenuItem("Open Config", "Open the config file")
	go func() {
		for {
			<-mOpenConfig.ClickedCh
			openFile(configFilePath)
		}
	}()

	mReload := systray.AddMenuItem("Reload", "Reload the config file")
	go func() {
		<-mReload.ClickedCh
		systray.ResetMenu()
		createMenu()
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

//go:embed icon.ico
var iconFile embed.FS

func getIcon(iconName string) []byte {
	data, _ := iconFile.ReadFile(iconName)
	return data
}

func onReady() {
	systray.SetIcon(getIcon("icon.ico"))
	systray.SetTitle("Quick Access")
	systray.SetTooltip("Quick Access")

	createMenu()
}

func onExit() {
	println("Exiting...")
}
