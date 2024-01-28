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
		ViewHandler: &viewHandler,
		OptionsList: &views.List{ListItems: []*views.ListItem{
			{Name: "Install BepInEx & Mods", Action: "install-both", Selected: true},
			{Name: "Only Install BepInEx", Action: "install-bepin", Selected: false},
			{Name: "Update Mods", Action: "install-updates", Selected: false},
		},
			Index: 0,
		},
	}

	homeView := handlers.View{
		Id:           "home",
		ViewRenderer: homeRenderer,
	}

	installRenderer := views.InstallView{}

	installView := handlers.View{
		Id:           "install",
		ViewRenderer: installRenderer,
	}

	viewHandler.AddView(homeView)
	viewHandler.AddView(installView)

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
