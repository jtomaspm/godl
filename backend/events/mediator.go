package events

import (
	"godl/backend/model"
)

type Mediator struct {
	handlers map[rune]*model.Action
}

func NewMediator() *Mediator {
	return &Mediator{
		handlers: make(map[rune]*model.Action),
	}
}

func (m *Mediator) Handle(event rune) {
	handler := m.handlers[event]
	if handler != nil {
		handler.Execute()
	}
}

func (m *Mediator) AddHandler(handler *model.Action) {
	m.handlers[handler.Hotkey] = handler
}

func (m *Mediator) AddHandlers(handlers []model.Action) {
	for i := range len(handlers) {
		m.AddHandler(&handlers[i])
	}
}
