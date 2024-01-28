package main

import (
	"fmt"
	"os"

	"github.com/JGugino/bepinstall/handlers"
	"github.com/JGugino/bepinstall/model"
	"github.com/JGugino/bepinstall/views"
	tea "github.com/charmbracelet/bubbletea"
)

var bepinVersionHandler handlers.BepinVersionHandler

func main() {
	configHandler := handlers.MustInitConfigHandler()

	viewHandler := handlers.InitViewHandler()

	bepinVersionHandler = handlers.InitBepinHandler()

	homeRenderer := views.HomeView{
		OptionsList: &views.List{ListItems: []*views.ListItem{
			{Name: "Install BepInEx & Mods", Selected: true},
			{Name: "Only Install BepInEx", Selected: false},
			{Name: "Update Mods", Selected: false},
		},
			Index: 0,
		},
	}

	homeView := handlers.View{
		Id:           "home",
		ViewRenderer: homeRenderer,
		ViewDisplay:  homeRenderer.Show(),
	}

	viewHandler.AddView(homeView)

	viewHandler.SetView("home")

	program := tea.NewProgram(model.InstallerModel{
		Id:            "installer-model-0.0.1",
		ConfigHandler: &configHandler,
		BepinHandler:  &bepinVersionHandler,
		ViewHandler:   &viewHandler,
	})

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
