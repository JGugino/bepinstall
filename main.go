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

	var homeView = handlers.View{
		Id:          "home",
		ViewDisplay: views.ShowHome(),
	}

	viewHandler.AddView(homeView)

	viewHandler.SetView("home")

	program := tea.NewProgram(model.InstallerModel{
		Id:            "installer-model-0.0.1",
		ConfigHandler: &configHandler,
		ViewHandler:   &viewHandler,
	})

	program.SetWindowTitle(fmt.Sprintf("BepInstaller v%s", configHandler.InstallerVersion))

	bepinVersionHandler = handlers.InitBepinHandler()

	//bepinVersionHandler.InstallBepinEx("5.4.22-win", configHandler, "testing")

	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
