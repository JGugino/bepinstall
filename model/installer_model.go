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

	HomeList        *handlers.List
	BepinVersions   *handlers.List
	GameDirectories *handlers.List

	CurrentView       string
	PossibleViews     []string
	primaryAction     string
	selectedVersion   string
	selectedDirectory string
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
		case "up", "w":
			if im.CurrentView == im.PossibleViews[0] {
				im.HomeList.MoveDown()
			}
			return im, nil
		case "down", "s":
			if im.CurrentView == im.PossibleViews[0] {
				im.HomeList.MoveDown()
			}
			return im, nil
		case "enter":
			if im.CurrentView == im.PossibleViews[0] {
				im.HomeList.SelectItem()
				im.primaryAction = im.HomeList.SelectedItem.Action
				if im.primaryAction == "install-bepin" {
					im.CurrentView = im.PossibleViews[1]
				} else if im.primaryAction == "install-mods" {
					im.CurrentView = im.PossibleViews[1]
				}
			} else if im.CurrentView == im.PossibleViews[1] {
				if im.primaryAction == "install-bepin" {
					im.BepinVersions.SelectItem()
					im.selectedVersion = im.BepinVersions.SelectedItem.Action
				}
			}

			return im, nil
		}
	}

	return im, nil
}

func (im InstallerModel) View() string {

	displayString := fmt.Sprintf("BepInstall - v%s\n\n", im.ConfigHandler.InstallerVersion)

	if im.CurrentView == im.PossibleViews[0] {
		displayString += "What would you like to do?\n\n"
		displayString += im.HomeList.RenderList()
	} else if im.CurrentView == im.PossibleViews[1] {
		displayString += "Choose a BepInEx Version\n\n"
		displayString += im.BepinVersions.RenderList()
	} else if im.CurrentView == im.PossibleViews[2] {
		displayString += "Install"
	}

	displayString += "\npress ctrl+c to quit.\n"

	return displayString
}
