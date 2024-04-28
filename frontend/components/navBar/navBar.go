package navBar

import (
	"godl/backend/model"
	"godl/frontend/components"
	"log"

	"github.com/rivo/tview"
)

type NavBarComponent struct {
	Component     *tview.Flex
	tabs          *tview.Flex
	margin        *tview.TextView
	maxTabs       int
	active        string
	tabComponents []components.TabComponent
	router        *Router
}

func NewNavBarComponent() *NavBarComponent {

	bar := tview.NewFlex()
	widgets := tview.NewFlex()

	barComponent := NavBarComponent{
		Component:     bar,
		margin:        components.PlaceholderTab().Component,
		tabs:          widgets,
		tabComponents: []components.TabComponent{},
		maxTabs:       12,
		active:        "",
		router:        NewRouter(),
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
	return []model.Action{}
}

func (b *NavBarComponent) AddTab(tab components.TabComponent) {
	b.tabs.AddItem(tab.Component, 0, 1, false)
	b.tabComponents = append(b.tabComponents, tab)
	b.resizeComponent()
	b.refreshActive()
}

func (b *NavBarComponent) refreshActive() {
	log.Println("Searching " + b.active)
	for i := 0; i < b.tabs.GetItemCount(); i++ {
		tab := b.tabs.GetItem(i).(*tview.TextView)
		if tab.GetText(true) == b.active {
			b.tabComponents[i].SetActive()
			log.Println("Found! " + b.active)
		} else {
			b.tabComponents[i].SetInactive()
			log.Println("Not found " + tab.GetText(true))
		}
	}
}

func (b *NavBarComponent) SetActive(tab string) {
	b.active = tab
	for i := 0; i < b.tabs.GetItemCount(); i++ {
		tab := b.tabs.GetItem(i).(*tview.TextView)
		if tab.GetText(true) == b.active {
			b.tabComponents[i].Activate()
		}
	}
	b.refreshActive()
}
