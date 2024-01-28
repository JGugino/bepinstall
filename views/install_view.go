package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type InstallView struct {
}

func (v InstallView) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return tea.Quit
		}
	}

	return nil
}

func (v InstallView) Show() string {
	installString := "\nInstalling the stuffs...\n\n"

	return installString
}
