package model

import (
	"fmt"

	"github.com/JGugino/bepinstall/handlers"
	tea "github.com/charmbracelet/bubbletea"
)

type InstallerModel struct {
	Id            string
	ConfigHandler *handlers.ConfigHandler
	ViewHandler   *handlers.ViewHandler
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

	displayString := fmt.Sprintf("BepInstall - v%s\n", im.ConfigHandler.InstallerVersion)

	displayString += im.ViewHandler.ShowView()

	displayString += "press ctrl+c to quit.\n"

	return displayString
}
