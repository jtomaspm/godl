package main

import "godl/frontend/app"

func main() {
	ac := app.NewAppContext()
	if err := ac.Run(); err != nil {
		panic(err)
	}
}
