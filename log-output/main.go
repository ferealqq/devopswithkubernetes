package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)


func randString(length int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length)
    rand.Read(b)
    return fmt.Sprintf("%x", b)[:length]
}

func main(){
	id := randString(10)
	for {
		log.Println(id)
		time.Sleep(5 * time.Second)
	}
}