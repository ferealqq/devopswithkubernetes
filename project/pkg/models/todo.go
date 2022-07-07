package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name" gorm:"size:140"`
	IsCompleted bool `json:"isCompleted" mapstructure:"isCompleted"`
}


func CreateTodo(n string) *Todo {
	return &Todo{
		Name: n,
		IsCompleted: false,
	}
}