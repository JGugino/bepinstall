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
	CurrentSubView    string
	PossibleSubViews  []string
	primaryAction     string
	selectedVersion   string
	selectedDirectory string
	installStarted    bool
	installFinished   bool
}

func (im InstallerModel) Init() tea.Cmd {
	return nil
}

func (im InstallerModel) resetListIndexes() {
	im.HomeList.ListItems[im.HomeList.Index].Selected = false
	im.HomeList.Index = 0
	im.HomeList.ListItems[im.HomeList.Index].Selected = true

	im.BepinVersions.ListItems[im.BepinVersions.Index].Selected = false
	im.BepinVersions.Index = 0
	im.BepinVersions.ListItems[im.BepinVersions.Index].Selected = true

	im.GameDirectories.ListItems[im.GameDirectories.Index].Selected = false
	im.GameDirectories.Index = 0
	im.GameDirectories.ListItems[im.GameDirectories.Index].Selected = true
}

func (im InstallerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if im.CurrentView == im.PossibleViews[1] {
				if im.CurrentSubView == im.PossibleSubViews[0] {
					im.CurrentView = im.PossibleViews[0]
				} else if im.CurrentSubView == im.PossibleSubViews[1] {
					if im.primaryAction == "install-bepin" {
						im.CurrentSubView = im.PossibleSubViews[0]
					} else {
						im.CurrentView = im.PossibleViews[0]
					}
				}
				im.resetListIndexes()
			} else if im.CurrentView == im.PossibleViews[2] {
				im.CurrentView = im.PossibleViews[1]
				im.resetListIndexes()
			}
			return im, nil
		case "ctrl+c":
			return im, tea.Quit
		case "up", "w":
			if im.CurrentView == im.PossibleViews[0] {
				im.HomeList.MoveUp()
			} else if im.CurrentView == im.PossibleViews[1] {
				if im.CurrentSubView == im.PossibleSubViews[0] {
					im.BepinVersions.MoveUp()
				} else if im.CurrentSubView == im.PossibleSubViews[1] {
					im.GameDirectories.MoveUp()
				}

			}
			return im, nil
		case "down", "s":
			if im.CurrentView == im.PossibleViews[0] {
				im.HomeList.MoveDown()
			} else if im.CurrentView == im.PossibleViews[1] {
				if im.CurrentSubView == im.PossibleSubViews[0] {
					im.BepinVersions.MoveDown()
				} else if im.CurrentSubView == im.PossibleSubViews[1] {
					im.GameDirectories.MoveDown()
				}
			}
			return im, nil
		case "enter":
			//Home View
			if im.CurrentView == im.PossibleViews[0] {
				im.HomeList.SelectItem()
				im.primaryAction = im.HomeList.SelectedItem.Action
				if im.primaryAction == "install-bepin" {
					//Sets current main view to select
					im.CurrentView = im.PossibleViews[1]

					//Sets current sub view to bepin select
					im.CurrentSubView = im.PossibleSubViews[0]
				} else if im.primaryAction == "install-mods" {
					im.CurrentView = im.PossibleViews[1]
					im.CurrentSubView = im.PossibleSubViews[1]
				}
				//Select View
			} else if im.CurrentView == im.PossibleViews[1] {
				if im.CurrentSubView == im.PossibleSubViews[0] {
					im.BepinVersions.SelectItem()
					im.selectedVersion = im.BepinVersions.SelectedItem.Action
					im.CurrentSubView = im.PossibleSubViews[1]
				} else if im.CurrentSubView == im.PossibleSubViews[1] {
					im.GameDirectories.SelectItem()
					im.selectedDirectory = im.GameDirectories.SelectedItem.Action
					im.CurrentView = im.PossibleViews[2]
				}
			} else if im.CurrentView == im.PossibleViews[2] {
				if im.primaryAction == "install-bepin" {
					im.installStarted = true
					im.installFinished = false
					err := im.BepinHandler.InstallBepinEx(im.selectedVersion, *im.ConfigHandler, im.selectedDirectory)

					if err != nil {
						panic(err)
					}

					im.installFinished = true
					im.installStarted = false
				} else if im.primaryAction == "install-mods" {
					return im, nil
				}
			}

			return im, nil
		}
	}

	return im, nil
}

func (im InstallerModel) View() string {

	displayString := fmt.Sprintf("BepInstall - v%s\n\n", im.ConfigHandler.InstallerVersion)

	//Home View
	if im.CurrentView == im.PossibleViews[0] {
		displayString += "What would you like to do?\n\n"
		displayString += im.HomeList.RenderList()

		//Select View
	} else if im.CurrentView == im.PossibleViews[1] {
		//Bepin Select
		if im.CurrentSubView == im.PossibleSubViews[0] {
			displayString += "Choose a BepInEx Version:\n\n"
			displayString += im.BepinVersions.RenderList()

			//Directory Select
		} else if im.CurrentSubView == im.PossibleSubViews[1] {
			displayString += "Choose where to install:\n\n"
			displayString += im.GameDirectories.RenderList()
		}

		//Install View
	} else if im.CurrentView == im.PossibleViews[2] {
		if im.primaryAction == "install-bepin" {
			if !im.installStarted && !im.installFinished {
				displayString += fmt.Sprintf("You're about to install ( %s ) in \"%s\"\n", im.BepinHandler.BepinDownloadLinks[im.selectedVersion].Name, im.ConfigHandler.GameDirectories[im.selectedDirectory])
				displayString += "Press 'enter' to continue\n"
			}
			if im.installStarted {
				displayString += fmt.Sprintf("Installing ( %s )...\n", im.BepinHandler.BepinDownloadLinks[im.selectedVersion].Name)
			}
			if im.installFinished {
				displayString += fmt.Sprintf("%s has been installed...\n", im.BepinHandler.BepinDownloadLinks[im.selectedVersion].Name)
			}
		} else if im.primaryAction == "install-mods" {
			displayString += fmt.Sprintf("You're about to install mods in %s\n", im.ConfigHandler.GameDirectories[im.selectedDirectory])
			displayString += "Press 'enter' to continue\n"
		}

	}

	//Show back text if not on home view, otherwise just show quit text
	if im.CurrentView != im.PossibleViews[0] {
		displayString += "\npress 'esc' to go back or 'ctrl+c' to quit.\n"
	} else {
		displayString += "\npress ctrl+c to quit.\n"
	}

	return displayString
}
