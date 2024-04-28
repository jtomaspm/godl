package components

import (
	"errors"
	"fmt"
	"godl/backend/model"
	"godl/frontend/router"

	"github.com/rivo/tview"
)

type NavBarComponent struct {
	Component  *tview.Flex
	tabs       *tview.Flex
	margin     *tview.TextView
	maxTabs    int
	currentTab int
	router     *router.Router
}

func NewNavBarComponent(r *router.Router) *NavBarComponent {
	barComponent := NavBarComponent{
		Component:  tview.NewFlex(),
		margin:     PlaceholderTab().Component,
		tabs:       tview.NewFlex(),
		maxTabs:    10,
		currentTab: -1,
		router:     r,
	}
	barComponent.resizeComponent()
	return &barComponent
}

func (b *NavBarComponent) updateCurrent(curr int) {
	if curr < 0 || curr >= b.tabs.GetItemCount() {
		return
	}
	b.currentTab = curr
	for i := 0; i < b.tabs.GetItemCount(); i++ {
		tab := b.tabs.GetItem(i).(*tview.TextView)
		if i == b.currentTab {
			SetTabActive(tab)
		} else {
			SetTabInactive(tab)
		}
	}
}

func (b *NavBarComponent) resizeComponent() {
	b.Component.Clear().
		AddItem(b.tabs, 0, b.tabs.GetItemCount(), false).
		AddItem(b.margin, 0, b.maxTabs-b.tabs.GetItemCount(), false)
}

func (b *NavBarComponent) GetActions() []model.Action {
	// Manual tab switch
	actions := []model.Action{}
	for i := 1; i <= 9; i++ {
		actions = append(actions, model.Action{
			DisplayTxt: fmt.Sprintf("Tab %d", i),
			Hotkey:     rune(i),
			Execute: func() {
				b.SetCurrent(i - 1)
				b.DrawCurrent()
			},
		})
	}
	actions = append(actions, model.Action{
		DisplayTxt: "Tab 10",
		Hotkey:     '0',
		Execute: func() {
			b.SetCurrent(9)
			b.DrawCurrent()
		},
	})

	// Dynamic tab switch
	actions = append(actions, model.Action{
		DisplayTxt: "Next Tab",
		Hotkey:     '.',
		Execute: func() {
			b.NextTab()
		},
	})
	actions = append(actions, model.Action{
		DisplayTxt: "Previous Tab",
		Hotkey:     ',',
		Execute: func() {
			b.PreviousTab()
		},
	})
	return actions
}

func (b *NavBarComponent) AddTab(label string, callback func()) error {
	tab := NewTabComponent(label, callback)
	if b.tabs.GetItemCount() >= b.maxTabs {
		return errors.New("max tabs reached")
	}
	b.tabs.AddItem(tab.Component, 0, 1, false)
	b.router.AddRoute(tab.label, tab.callback)
	b.resizeComponent()

	return nil
}

func (b *NavBarComponent) SetCurrent(curr int) error {
	if curr < 0 || curr >= b.tabs.GetItemCount() {
		return fmt.Errorf("invalid current value: %d", curr)
	}
	b.updateCurrent(curr)
	return nil
}

func (b *NavBarComponent) DrawCurrent() error {
	if b.currentTab < 0 || b.currentTab >= b.tabs.GetItemCount() {
		return fmt.Errorf("invalid current tab: %d", b.currentTab)
	}
	label := b.tabs.GetItem(b.currentTab).(*tview.TextView).GetText(true)
	b.router.Draw(label)
	return nil
}

func (b *NavBarComponent) NextTab() {
	calcNext := func(b *NavBarComponent) {
		totalTabs := b.tabs.GetItemCount()
		if totalTabs == 0 {
			return
		}
		if b.currentTab >= totalTabs-1 {
			b.updateCurrent(0)
			return
		}
		b.updateCurrent(b.currentTab + 1)
	}
	calcNext(b)
	b.DrawCurrent()
}

func (b *NavBarComponent) PreviousTab() {
	calcPrev := func(b *NavBarComponent) {
		totalTabs := b.tabs.GetItemCount()
		if totalTabs == 0 {
			return
		}
		if b.currentTab == 0 {
			b.updateCurrent(totalTabs - 1)
			return
		}
		b.updateCurrent(b.currentTab - 1)
	}
	calcPrev(b)
	b.DrawCurrent()
}
