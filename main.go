package main

import (
	"fmt"
	"os"

	"github.com/JGugino/bepinstall/handlers"
	tea "github.com/charmbracelet/bubbletea"
)

type InstallerModel struct {
	id string
}

var bepinVersionHandler handlers.BepinVersionHandler

func main() {
	configHandler := handlers.InitConfigHandler()

	fmt.Println(configHandler)

	program := tea.NewProgram(initModel())

	bepinVersionHandler = handlers.InitBepinHandler()

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initModel() InstallerModel {
	return InstallerModel{
		id: "installer-model-0.0.1",
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

	return displayString
}
