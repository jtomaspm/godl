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

func (b *NavBarComponent) resizeComponent() {
	b.Component.Clear().
		AddItem(b.tabs, 0, b.tabs.GetItemCount(), false).
		AddItem(b.margin, 0, b.maxTabs-b.tabs.GetItemCount(), false)
}

func (b *NavBarComponent) GetActions() []model.Action {
	return []model.Action{
		{
			DisplayTxt: "Tab 1",
			Hotkey:     '1',
			Execute: func() {
				b.SetCurrent(1)
				b.DrawCurrent()
			},
		},
	}
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
	b.currentTab = curr
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
			b.currentTab = 0
			return
		}
		b.currentTab += 1
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
			b.currentTab = totalTabs - 1
			return
		}
		b.currentTab -= 1
	}
	calcPrev(b)
	b.DrawCurrent()
}
