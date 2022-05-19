package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

var (
	port    = os.Getenv("PORT")
	version = "0.01"
)


func main(){
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	startupMessage := "===> Starting server (v" + version + ")"
	startupMessage = startupMessage + " on port " + port
	log.Println(startupMessage)

	n.Run("localhost:" + port)
}