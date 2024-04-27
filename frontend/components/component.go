package components

import "godl/backend/model"

type Component interface {
	GetActions() []model.Action
}
