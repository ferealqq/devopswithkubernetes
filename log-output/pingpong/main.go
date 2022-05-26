package main

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const cStr = "Ping / Pongs: "
const file = "files/log.txt"

func readFile() (string, error) {
	if b, err := os.ReadFile(file); err == nil {
		return bytes.NewBuffer(b).String(), nil
	} else {
		return "", err
	}
}

func writeCount() (string, error) {
	if str, err := readFile(); err == nil {
		arr := strings.Split(str,cStr)
		var s string
		if len(arr) > 1 { 
			cs := arr[len(arr)-1]
			if c, e := strconv.Atoi(cs); e == nil {
				s = arr[0]+cStr+strconv.Itoa(c+1)
			} else {
				return "", e
			}
		} else {
			s = cStr+"1"
		}
		return s, os.WriteFile(file, []byte(s), 0644)
	} else if os.IsNotExist(err) {
		s := cStr+"1"
		return s, os.WriteFile(file, []byte(s), 0644)
	} else {
		return "", err
	}
}


func main(){
	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request){
		if c, e := writeCount(); e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}else {
			http.ServeContent(w,r,"kontsaa",time.Now(),bytes.NewReader([]byte(c)))
		}
	})	

	http.ListenAndServe(":3001", nil)
}