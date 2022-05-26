package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"
)


func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// miten vitussä tää voi lukee tota log.txt filuu ku se on files/log.txt kohas?
		if b, err := os.ReadFile("files/log.txt"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}else{
			http.ServeContent(w,r,"",time.Now(),bytes.NewReader(b))
		}
	})	

	log.Println(http.ListenAndServe(":3003", nil))
}