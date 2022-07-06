package main

import (
	"log"
	"net/http"
	"project-todo/pkg/db"
	"project-todo/pkg/models"
)


func main(){
	conn := db.Conn()

	if err := conn.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	if res, err := client.Get("https://en.wikipedia.org/wiki/Special:Random"); err != nil {
		log.Fatal(err)
	}else{
		if link, e := res.Location(); e != nil {
			log.Fatal(err)
		}else{
			conn.Create(models.CreateTodo("Read "+link.String()))
		}
	}
}