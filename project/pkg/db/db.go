package db

import (
	"project-todo/pkg/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Conn() *gorm.DB {
	if db == nil {
		var err error
		pq := postgres.Open(util.GetEnv("DB_DSN", "host=localhost user=this password=this sslmode=disable"))
		if db, err = gorm.Open(pq, &gorm.Config{}); err == nil {
			return db
		}else{
			panic(err)
		}
	}
	return db
}

func Close() {
	if d, err := db.DB(); err != nil {
		panic(err)
	}else{
		d.Close()
	}
}
