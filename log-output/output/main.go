package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const file = "files/log.txt"
func randString(length int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length)
    rand.Read(b)
    return fmt.Sprintf("%x", b)[:length]
}

func main(){
	// ping pong count start
	const cStr = "Ping / Pongs: "
	id := randString(10)
	// don't catch error because it may already exists and if it does we do not care that it failed
	os.Mkdir("files",0755);
	for {
		cur := time.Now().String()+ " : "+ id+"\n"
		if b,err := os.ReadFile(file); err == nil {
			s := bytes.NewBuffer(b).String()
			arr := strings.Split(s, cStr)
			var out string
			if len(arr) > 1 { 
				out = cur+cStr+arr[1]
			}else if strings.Contains(s, cStr){
				out = cur+cStr+arr[0]
			}else{
				out = cur+cStr+"0"
			}
			fmt.Println(out)
			if e := os.WriteFile(file,[]byte(out), 0644); e != nil {
				log.Fatal(err.Error())
			}
		} else if os.IsNotExist(err) {
			out := cur+cStr+"0"
			fmt.Println(out)
			if e := os.WriteFile(file,[]byte(out), 0644); e != nil {
				log.Fatal(err.Error())
			}
		} else {
			log.Println(err)
			log.Fatal(err.Error())
		}
		time.Sleep(5 * time.Second)
	}	
}