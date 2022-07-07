package main

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}


const cStr = "Ping / Pongs: "

var sConn = getEnv("CONNECTION","host=localhost user=this password=this sslmode=disable")

var db *sql.DB

func increase() (int,error){
	if db != nil {
		var c int 
		err := db.QueryRow("update counter set n = n + 1 where id = 1 returning n").Scan(&c);
		if err != nil{
			return 0, err 
		}

		return c, nil
	}
	// fixme return a error
	return 0,nil
}

func count() (int, error){
	if db != nil {
		var c int 
		err := db.QueryRow("select n from counter where id = 1 limit 1;").Scan(&c);
		if err != nil{
			return 0, err 
		}

		return c, nil
	}
	// fixme return a error
	return 0,nil
}

func initCounter() {
	q := "insert into counter (id,n)"+
		 "select 1, 0 "+
		 "where not exists ( "+
			"select id from counter where id = 1"+
		 ");"	

	if _,err := db.Query(q); err != nil {
		log.Fatal(err)
	}
}

func main(){
	conn, err := sql.Open("postgres",sConn)
	if err == nil {
		db = conn
		initCounter()
	}else{
		panic(err)
	}

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request){
		if count, err := increase(); err == nil {
			c := cStr+strconv.Itoa(count)
	
			http.ServeContent(w,r,"kontsaa", time.Now(), bytes.NewReader([]byte(c)))
		}else{
			log.Fatal(err)
			http.ServeContent(w,r,"k",time.Now(), bytes.NewReader([]byte("something went wrong, try again later")))
		}
	})
	
	http.HandleFunc("/pingpong/count", func(w http.ResponseWriter, r *http.Request){
		if count, err := count(); err == nil {
			http.ServeContent(w,r,"count", time.Now(), bytes.NewReader([]byte(strconv.Itoa(count))))
		}else{
			log.Fatal(err)
			http.ServeContent(w,r,"count", time.Now(), bytes.NewReader([]byte(strconv.Itoa(0))))
		}
	});

	log.Fatal(http.ListenAndServe(":"+getEnv("PORT", "3001"), nil))
}