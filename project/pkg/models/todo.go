package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	IsCompleted bool `json:"isCompleted" mapstructure:"isCompleted"`
}


func CreateTodo(n string) *Todo {
	return &Todo{
		Name: n,
		IsCompleted: false,
	}
}