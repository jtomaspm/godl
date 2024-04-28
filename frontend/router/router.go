package router

type route string

type Router struct {
	routes map[route]func()
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[route]func()),
	}
}

func (m *Router) Draw(r route) {
	m.routes[r]()
}

func (m *Router) AddRoute(r route, callback func()) {
	m.routes[r] = callback
}

func (m *Router) GetResolution(r route) func() {
	return m.routes[r]
}
