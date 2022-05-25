package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func getEnv(key string, fallback string) string {
	if v, s := os.LookupEnv(key); s {
		return v
	}
	return fallback
}

var (
	port    = getEnv("PORT", "3000")
	version = "0.03"
)


func main(){
	router := gin.Default()
	router.StaticFile("/","static/index.html")
	router.StaticFile("/wasm_exec.js","static/wasm_exec.js")
	router.StaticFile("/app.wasm", "static/app.wasm")
	gin.SetMode(gin.ReleaseMode)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	startupMessage := "===> Starting server (v" + version + ")"
	startupMessage = startupMessage + " on port " + port
	log.Println(startupMessage)
	
	n.Run(":"+port)
}