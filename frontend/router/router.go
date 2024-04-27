package router

import "godl/frontend/screens"

type route string

type Router struct {
	routes map[route]screens.Screen
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[route]screens.Screen),
	}
}

func (m *Router) Draw(r route) {
	m.routes[r].Draw()
}

func (m *Router) AddRoute(r route, s screens.Screen) {
	m.routes[r] = s
}
