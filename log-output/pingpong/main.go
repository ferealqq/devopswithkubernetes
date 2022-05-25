package main

import (
	"bytes"
	"net/http"
	"strconv"
	"time"
)


func main(){
	count := 0
	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request){
		count++
		c := "Count: "+strconv.Itoa(count)
		http.ServeContent(w,r,"kontsaa",time.Now(),bytes.NewReader([]byte(c)))
	})

	http.ListenAndServe(":3001", nil)
}