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
		HomeList: handlers.CreateNewList(0, []*handlers.ListItem{
			{Name: "Install BepInEx", Action: "install-bepin", Selected: true},
			{Name: "Install Mods", Action: "install-mods", Selected: false},
		}),
		BepinVersions: handlers.CreateNewList(0, []*handlers.ListItem{
			{Name: bepinVersionHandler.BepinDownloadLinks["5.4.22-win"].Name, Action: bepinVersionHandler.BepinDownloadLinks["5.4.22-win"].Id, Selected: true},
			{Name: bepinVersionHandler.BepinDownloadLinks["5.4.22-unix"].Name, Action: bepinVersionHandler.BepinDownloadLinks["5.4.22-unix"].Id, Selected: false},
		}),
	}

	program := tea.NewProgram(installerModel)

	program.SetWindowTitle(fmt.Sprintf("BepInstaller v%s", configHandler.InstallerVersion))

	// err := bepinVersionHandler.InstallBepinEx("5.4.22-win", configHandler, "testing")

	// if err != nil {
	// 	panic(err)
	// }

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
