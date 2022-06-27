package app

import (
	"encoding/json"
	"log"
	"net/http"
	"project-todo/pkg/models"
	"project-todo/pkg/util"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/mitchellh/mapstructure"
)

func main() {
	vecty.SetTitle("DevopsWithKubernetes")
	vecty.RenderBody(&App{})
}

var (
	service = util.GetEnv("SERVICE", "localhost")
	port = util.GetEnv("SERVICE_PORT", "4000")
)

var api = "http://"+service+":"+port

type App struct {
	vecty.Core
}

func fetchTodos() (*[]models.Todo, error) {
	if res, err := http.Get(api); err != nil {
		return nil, err
	}else{
		var data map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&data)
		var todos []models.Todo
		if data != nil && err == nil {
			return &todos, mapstructure.Decode(data, &todos)
		}
		return nil, err

	}
}

func (p *App) Render() vecty.ComponentOrHTML {
	todoInput := "todo-input-id"
	todos, err := fetchTodos()
	if err != nil {
		log.Println(err.Error())
		return vecty.Text("Something went wrong try again");
	}
	var items vecty.List
	for _, item := range *todos {
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