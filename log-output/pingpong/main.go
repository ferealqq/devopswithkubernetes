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

// Depricated
func readFile() (string, error) {
	if b, err := os.ReadFile(file); err == nil {
		return bytes.NewBuffer(b).String(), nil
	} else {
		return "", err
	}
}
// Depricated
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


func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func main(){
	count := 0
	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request){
		count++

		c := cStr+strconv.Itoa(count)
	
		http.ServeContent(w,r,"kontsaa", time.Now(), bytes.NewReader([]byte(c)))
	})
	
	http.HandleFunc("/pingpong/count", func(w http.ResponseWriter, r *http.Request){
		http.ServeContent(w,r,"count", time.Now(), bytes.NewReader([]byte(strconv.Itoa(count))))
	});

	http.ListenAndServe(":"+getEnv("PORT", "3001"), nil)
}