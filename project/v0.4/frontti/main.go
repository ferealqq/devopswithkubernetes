package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func main() {
	vecty.SetTitle("DevopsWithKubernetes")
	vecty.RenderBody(&App{})
}

type App struct {
	vecty.Core
}

func (p *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Style("float", "center"),
			),
			elem.Heading1(
				vecty.Text("Hello DevOps with Kubernetes"),
			),
		),
	)
}