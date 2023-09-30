package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

var restart = false

func onReady() {
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

	mRestart := systray.AddMenuItem("Restart", "Restart the app")
	go func() {
		<-mRestart.ClickedCh
		restart = true
		systray.Quit()
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	systray.SetIcon(getIcon("icon.ico"))
	systray.SetTooltip("Quick Access")
}

func onExit() {
	println("Exiting...")

	if restart {
		executablePath, err := os.Executable()
		if err != nil {
			fmt.Println("Failed to get executable path:", err)
			return
		}

		cmd := exec.Command(executablePath)
		err = cmd.Start()
		if err != nil {
			fmt.Println("Failed to restart executable:", err)
			return
		}

		fmt.Println("Restarted the executable.")
	}
}
