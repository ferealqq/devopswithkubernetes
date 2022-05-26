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
	todoInput := "todo-input-id"
	todoItems := []string{"Exercise 1.12", "Exrcise 1.13"}
	var items vecty.List
	for _, item := range todoItems {
		items = append(items, 
			elem.ListItem(
				vecty.Markup(
					vecty.Class("list-group-item"),
				),
				vecty.Text(item),
			),
		)
	} 
	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css")
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
			),
			elem.Heading1(
				vecty.Text("Hello DevOps with Kubernetes"),
			),
			elem.Image(
				vecty.Markup(
					vecty.Style("width", "25%"),
					vecty.Style("height", "auto"),
					vecty.Attribute("src", "/images/today.png"),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("container"),
				),
				elem.Label(
					vecty.Markup(
						vecty.Class("form-label"),
						vecty.Attribute("for", todoInput),
					),

					vecty.Text("TODO"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("row"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("col-8"),
						),
						elem.Input(
							vecty.Markup(
								vecty.Class("form-control"),
								vecty.Attribute("type", "text"),
								vecty.Attribute("id", todoInput),
								vecty.Attribute("maxlength", 140),
							),
						),	
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("col-4"),
						),
						elem.Button(
							vecty.Markup(
								vecty.Class("btn"),
								vecty.Class("btn-primary"),
							),
							vecty.Text("Save"),
						),	
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("col-8"),
					),
					
					elem.UnorderedList(
						vecty.Markup(
							vecty.Class("list-group"),
							vecty.Class("py-2"),
						),
						items,
					),
				),
			),
		),
	)
}