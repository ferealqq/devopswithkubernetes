package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"project-todo/pkg/util"

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

func saveTodayImage(){
	os.Mkdir("images", 0755)
	for {
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
		time.Sleep(24 * time.Hour)
	}
}

func main(){
	router := gin.Default()
	router.Use(util.CORSMiddleware())

	// this could and should be don with a cron job but we haven't figured out that part yet about kubernetes
	go saveTodayImage()
	router.StaticFile("/","static/index.html")
	router.StaticFile("/wasm_exec.js","static/wasm_exec.js")
	router.StaticFile("/app.wasm", "static/app.wasm")
	router.StaticFile("/images/today.png", "images/today.png")
	gin.SetMode(gin.ReleaseMode)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	startupMessage := "===> Starting server (v" + version + ")"
	startupMessage = startupMessage + " on port " + port
	log.Println(startupMessage)
	
	n.Run(":"+port)
}
