package handlers

import (
	"errors"
)

type ViewHandler struct {
	CurrentView   View
	possibleViews map[string]View
}

type View struct {
	Id          string
	ViewDisplay string
}

func InitViewHandler() ViewHandler {
	return ViewHandler{
		CurrentView:   View{},
		possibleViews: make(map[string]View, 4),
	}
}

func (h *ViewHandler) ShowView() string {
	returnString := ""

	returnString += h.CurrentView.ViewDisplay

	returnString += "\n"

	return returnString
}

func (h *ViewHandler) AddView(view View) error {
	if _, ok := h.possibleViews[view.Id]; !ok {
		h.possibleViews[view.Id] = view
		return nil
	}

	return errors.New("view-already-exists")
}

func (h ViewHandler) ChangeView(viewId string) error {
	return nil
}

func (h *ViewHandler) SetView(viewId string) error {
	if _, ok := h.possibleViews[viewId]; !ok {
		return errors.New("unknown-view")
	}

	h.CurrentView = h.possibleViews[viewId]

	return nil
}

func (h ViewHandler) GetPossibleViews() map[string]View {
	return h.possibleViews
}
