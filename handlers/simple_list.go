package handlers

import "fmt"

type List struct {
	ListItems    []*ListItem
	SelectedItem ListItem
	Index        int
}

type ListItem struct {
	Name     string
	Action   string
	Selected bool
}

func CreateNewList(startingIndex int, listItems []*ListItem) *List {
	listItems[0].Selected = true
	return &List{
		ListItems: listItems,
		Index:     startingIndex,
	}
}

func (l *List) MoveUp() {
	if l.Index <= 0 {
		l.ListItems[l.Index].Selected = false
		l.Index = 0
		l.ListItems[l.Index].Selected = true
	} else {
		l.ListItems[l.Index].Selected = false
		l.Index--
		l.ListItems[l.Index].Selected = true
	}
}

func (l *List) MoveDown() {
	if l.Index >= len(l.ListItems)-1 {
		l.ListItems[l.Index].Selected = false
		l.Index = len(l.ListItems) - 1
		l.ListItems[l.Index].Selected = true
	} else {
		l.ListItems[l.Index].Selected = false
		l.Index++
		l.ListItems[l.Index].Selected = true
	}
}

func (l *List) SelectItem() ListItem {
	l.SelectedItem = *l.ListItems[l.Index]
	return l.SelectedItem
}

func (l *List) RenderList() string {
	listDisplay := ""

	for _, v := range l.ListItems {
		if v.Selected {
			listDisplay += fmt.Sprintf("> %s\n", v.Name)
		} else {
			listDisplay += fmt.Sprintf("  %s\n", v.Name)
		}
	}

	return listDisplay
}
