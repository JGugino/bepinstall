package main

import (
	"fmt"
	"os"

	"github.com/JGugino/bepinstall/handlers"
	"github.com/JGugino/bepinstall/model"
	tea "github.com/charmbracelet/bubbletea"
)

var bepinVersionHandler handlers.BepinVersionHandler

func main() {
	configHandler := handlers.MustInitConfigHandler()

	bepinVersionHandler = handlers.InitBepinHandler()

	gameDirectories := make([]*handlers.ListItem, 0)

	for k := range configHandler.GameDirectories {
		gameDirectories = append(gameDirectories, &handlers.ListItem{Name: k, Action: k, Selected: false})
	}

	installerModel := &model.InstallerModel{
		Id:            "installer-model-0.0.1",
		ConfigHandler: &configHandler,
		BepinHandler:  &bepinVersionHandler,
		CurrentView:   "home",
		PossibleViews: []string{
			"home",
			"select",
			"install",
		},
		CurrentSubView: "bepin",
		PossibleSubViews: []string{
			"bepin",
			"mods",
		},
		HomeList: handlers.CreateNewList(0, []*handlers.ListItem{
			{Name: "Install BepInEx", Action: "install-bepin", Selected: true},
			{Name: "Install Mods", Action: "install-mods", Selected: false},
		}),
		BepinVersions: handlers.CreateNewList(0, []*handlers.ListItem{
			{Name: bepinVersionHandler.BepinDownloadLinks["5.4.22-win"].Name, Action: bepinVersionHandler.BepinDownloadLinks["5.4.22-win"].Id, Selected: false},
			{Name: bepinVersionHandler.BepinDownloadLinks["5.4.22-unix"].Name, Action: bepinVersionHandler.BepinDownloadLinks["5.4.22-unix"].Id, Selected: false},
		}),
		GameDirectories: handlers.CreateNewList(0, gameDirectories),
	}

	program := tea.NewProgram(installerModel)

	program.SetWindowTitle(fmt.Sprintf("BepInstaller v%s", configHandler.InstallerVersion))

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
