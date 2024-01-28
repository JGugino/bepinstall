package model

import (
	"fmt"

	"github.com/JGugino/bepinstall/handlers"
	tea "github.com/charmbracelet/bubbletea"
)

type InstallerModel struct {
	Id            string
	ConfigHandler *handlers.ConfigHandler
	BepinHandler  *handlers.BepinVersionHandler
	ViewHandler   *handlers.ViewHandler
}

func (im InstallerModel) Init() tea.Cmd {
	return nil
}

func (im InstallerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return im, im.ViewHandler.CurrentView.ViewRenderer.Update(msg)
}

func (im InstallerModel) View() string {

	displayString := fmt.Sprintf("BepInstall - v%s\n", im.ConfigHandler.InstallerVersion)

	displayString += im.ViewHandler.CurrentView.ViewRenderer.Show()

	displayString += "\npress ctrl+c to quit.\n"

	return displayString
}
