package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project-todo/pkg/models"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"

	"github.com/mitchellh/mapstructure"
)

func main() {
	vecty.SetTitle("DevopsWithKubernetes")
	vecty.RenderBody(&App{})
}

type App struct {
	vecty.Core
	newTodo string
	todos *[]models.Todo
}

func fetchTodos() (*[]models.Todo, error) {
	if res, err := http.Get("/api/todos"); err != nil {
		return nil, err
	}else{
		var data []map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&data)
		var todos []models.Todo
		if data != nil && err == nil {
			return &todos, mapstructure.Decode(data, &todos)
		}
		return nil, err
	}
}

func Post(path string,data interface{}) (*http.Response, error){
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(data); err != nil {
		return nil, err
	}
	return http.Post("api/"+path,"application/json",b)
}

func (a *App) postTodo(td *models.Todo){
	if resp, err := Post("todos",td); err != nil {
		println(err)
	}else{
		var data map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&data)
		var td models.Todo
		if data != nil && err == nil {
			mapstructure.Decode(data, &td)
			println("appending todo list: "+td.Name)
			*a.todos = append(*a.todos,td)
			println("rerender pls")
			vecty.Rerender(a)
		}else{
			println(err)
		}
	}
}

func (a *App) onPostTodo(_ *vecty.Event) {
	go a.postTodo(models.CreateTodo(a.newTodo))
	a.newTodo = ""
}

func (a *App) onTodoInput(event *vecty.Event){
	a.newTodo = event.Target.Get("value").String()
}

func (a *App) Mount(){
	Listen(func(){
		vecty.Rerender(a)
	})
}

func (a *App) fetchTodos() {
	// if todos has been fetched, do not fetch them again
	if a.todos != nil {
		return
	}
	if tds, err := fetchTodos(); err == nil {
		a.todos = tds
		Dispatch()
	}else{
		println("something went wrong")
	}
}

func (p *App) Render() vecty.ComponentOrHTML {
	todoInput := "todo-input-id"

	go p.fetchTodos()

	var items vecty.List
	if p.todos != nil {
		for _, item := range *p.todos {
			println(item.Name)
			items = append(items, 
				elem.ListItem(
					vecty.Markup(
						vecty.Class("list-group-item"),
					),
					vecty.Text(item.Name),
				),
			)
		} 
	}

	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css")
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
			),
			elem.Heading1(
				vecty.Text("Hello DevOps with Kubernetes!"),
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
								prop.Placeholder("What needs to be done? Nothing absolutely, nothing!"),
								prop.Autofocus(true),
								prop.Value(p.newTodo),
								event.Input(p.onTodoInput),
							),
						),	
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("col-4"),
							event.Click(p.onPostTodo),
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
