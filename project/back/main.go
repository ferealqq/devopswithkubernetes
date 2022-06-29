package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"project-todo/pkg/models"
	"project-todo/pkg/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

var (
	port    = util.GetEnv("PORT", "4000")
	version = "1.00"
)

var cImg = make(chan struct{})

func saveTodayImage(){
	os.Mkdir("images", 0755)
	tic := time.NewTicker(time.Hour*24)
	loop:
	for {
		select {
		case <-cImg:
			// stop saving images
			break loop
		case <-tic.C:
			if resp, err := http.Get("https://picsum.photos/1200"); err == nil {
				if f, e := os.Create("images/today.png"); e == nil {
					if _, e = io.Copy(f, resp.Body); e != nil {
						f.Close()
						resp.Body.Close()
						log.Fatal(e.Error())
					}

					f.Close()
					resp.Body.Close()
				}else{
					f.Close()
					resp.Body.Close()
					log.Fatal(e.Error())
				}
				
			}else{
				log.Fatal(err.Error())
			}
		}
	}
}

func main(){
	var todos []*models.Todo
	todos = append(todos, models.CreateTodo("Exercise 1.12"), models.CreateTodo("Exrcise 1.13"))
	router := gin.Default()
	router.Use(util.CORSMiddleware())
	// this could and should be don with a cron job but we haven't figured out that part yet about kubernetes
	go saveTodayImage()
	router.GET("/api/todos", func(c *gin.Context){
		c.JSON(200, todos)
	})
	router.POST("/api/todos", func(c *gin.Context){
		var todo *models.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(500, err.Error())
		}else{
			todos = append(todos, todo)
			c.JSON(200, todo)
		}
	})
	router.StaticFile("/images/today.png", "images/today.png")
	gin.SetMode(gin.ReleaseMode)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	startupMessage := "===> Starting API (v" + version + ")"
	startupMessage = startupMessage + " on port " + port
	log.Println(startupMessage)
	
	n.Run(":"+port)
}