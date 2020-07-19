package main

import (
	"geektest/app"
	"geektest/internal"
	"geektest/internal/routing"
)

func main() {
	app.SetRoutes(routing.GetRouter())
	internal.StartServer()
}