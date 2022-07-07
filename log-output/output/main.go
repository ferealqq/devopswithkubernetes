package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
		if e := os.WriteFile(file, []byte(cur), 0644); e != nil {
			log.Fatal(e.Error())
		}

		time.Sleep(5 * time.Second)
	}	
}