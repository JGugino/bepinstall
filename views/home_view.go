package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type HomeView struct {
	OptionsList *List
}

type List struct {
	ListItems []*ListItem
	Index     int
}

type ListItem struct {
	Name     string
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
