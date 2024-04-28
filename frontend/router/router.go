package router

type Router struct {
	routes map[string]func()
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func()),
	}
}

func (m *Router) Draw(r string) {
	m.routes[r]()
}

func (m *Router) AddRoute(r string, callback func()) {
	m.routes[r] = callback
}

func (m *Router) GetResolution(r string) func() {
	return m.routes[r]
}
