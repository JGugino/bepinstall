package main

import (
	"fmt"
	"os"

	"github.com/JGugino/bepinstall/handlers"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type InstallerModel struct {
	id               string
	progressBar      progress.Model
	currentPercent   float64
	progressBarShown bool
}

var bepinVersionHandler handlers.BepinVersionHandler

func main() {
	configHandler := handlers.MustInitConfigHandler()

	progressBar := progress.New(progress.WithColorProfile(termenv.Ascii))

	program := tea.NewProgram(InstallerModel{
		id:               "installer-model-0.0.1",
		progressBar:      progressBar,
		currentPercent:   0.2,
		progressBarShown: false,
	})

	bepinVersionHandler = handlers.InitBepinHandler()

	bepinVersionHandler.InstallBepinEx("5.4.22-win", configHandler, "testing")

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (im InstallerModel) Init() tea.Cmd {
	return nil
}

func (im InstallerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return im, tea.Quit
		}
		return im, nil
	}

	return im, nil
}

func (im InstallerModel) View() string {

	displayString := "BepInstall - v0.0.1\npress ctrl+c to quit.\n\n"

	if im.progressBarShown {
		displayString += im.progressBar.ViewAs(im.currentPercent) + "\n\n"
	}
	return displayString
}
