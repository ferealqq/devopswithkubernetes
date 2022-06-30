package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func main(){
	const cStr = "Ping / Pongs: "

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if b, err := os.ReadFile("files/log.txt"); err != nil {
			http.Error(w, "Reading: "+err.Error(), http.StatusInternalServerError)
			return
		}else{
			resp, err := http.Get("http://pingpong-svc:"+getEnv("PINGPONG_PORT","3001")+"/pingpong/count")
			if err == nil {
				var count int
				if e := json.NewDecoder(resp.Body).Decode(&count); e != nil {
					log.Fatal(e.Error())
				} 
				c := getEnv("MESSAGE", "Ei toimi")+"\n"+string(b)+cStr+strconv.Itoa(count)	
				http.ServeContent(w,r,"",time.Now(),bytes.NewReader([]byte(c)))
			}else{
				http.Error(w, "Fetching: "+err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
			}
		}
	})	

	log.Println(http.ListenAndServe(":3003", nil))
}