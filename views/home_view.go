package views

import (
	"fmt"

	"github.com/JGugino/bepinstall/handlers"
	tea "github.com/charmbracelet/bubbletea"
)

type HomeView struct {
	ViewHandler *handlers.ViewHandler
	OptionsList *List
}

type List struct {
	ListItems []*ListItem
	Index     int
}

type ListItem struct {
	Name     string
	Action   string
	Selected bool
}

func (v HomeView) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return tea.Quit
		case "up", "w":
			if v.OptionsList.Index <= 0 {
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = false
				v.OptionsList.Index = 0
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = true
			} else {
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = false
				v.OptionsList.Index--
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = true
			}
			return nil
		case "down", "s":
			if v.OptionsList.Index >= len(v.OptionsList.ListItems)-1 {
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = false
				v.OptionsList.Index = len(v.OptionsList.ListItems) - 1
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = true
			} else {
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = false
				v.OptionsList.Index++
				v.OptionsList.ListItems[v.OptionsList.Index].Selected = true
			}
			return nil
		case "enter":
			if v.OptionsList.ListItems[v.OptionsList.Index].Action == "install-both" {
				v.ViewHandler.SetView("install")
				fmt.Println("Install Both")
			} else if v.OptionsList.ListItems[v.OptionsList.Index].Action == "install-bepin" {
				v.ViewHandler.SetView("install")
				fmt.Println("Install Bepin")
			} else if v.OptionsList.ListItems[v.OptionsList.Index].Action == "install-updates" {
				v.ViewHandler.SetView("install")
				fmt.Println("Install Updates")
			}
			return nil
		}
		return nil
	}

	return nil
}

func (v HomeView) Show() string {
	homeString := "\nWhat would you like to do?\n\n"

	for _, v := range v.OptionsList.ListItems {
		if v.Selected {
			homeString += fmt.Sprintf("> %s\n", v.Name)
		} else {
			homeString += fmt.Sprintf("  %s\n", v.Name)
		}
	}

	return homeString
}
