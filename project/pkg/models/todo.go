package models

type Todo struct {
	Name string `json:"name"`
	IsCompleted bool `json:"isCompleted" mapstructure:"isCompleted"`
}


func CreateTodo(n string) *Todo {
	return &Todo{
		Name: n,
		IsCompleted: false,
	}
}