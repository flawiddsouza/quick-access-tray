package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

var restart = false

func onReady() {
	config, err := parseConfigYAML("config.yml")
	if err != nil {
		println("Failed to parse config.yml:", err)
		return
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
