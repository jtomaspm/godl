package events

import (
	"godl/frontend/model"
)

type Mediator struct {
	handlers map[rune]model.Action
}

func NewMediator() *Mediator {
	return &Mediator{
		handlers: make(map[rune]model.Action),
	}
}

func (m *Mediator) Handle(event rune) {
	m.handlers[event].Execute()
}

func (m *Mediator) Addhandler(handler model.Action) {
	m.handlers[handler.Hotkey] = handler
}
