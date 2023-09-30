package main

import (
	"github.com/getlantern/systray"
)

type Project struct {
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	projects, err := parseProjectsYAML()
	if err != nil {
		println("Failed to parse projects.yml:", err)
		return
	}

	for _, project := range projects {
		menuItem := systray.AddMenuItem(project.Label, project.Label)
		go func(command string) {
			for {
				select {
				case <-menuItem.ClickedCh:
					openProject(command)
				}
			}
		}(project.Command)
	}

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	systray.SetIcon(getIcon("icon.ico"))
	systray.SetTooltip("Projects")
}

func onExit() {
	println("Exiting...")
}
