package config

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	ConfServerPort  = "3000"
	ConfBaseURL     = "server.base_url"
	ConfCORSEnabled = "true"

	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "learn-event-sourcing"
	DB_HOST = "localhost"
	DB_PORT = "5432"
)

var db *gorm.DB
var dbOnce sync.Once

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		var conn string
		conn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			DB_HOST,
			DB_PORT,
			DB_USER,
			DB_PASS,
			DB_NAME,
		)

		db, err = gorm.Open("postgres", conn)
		if nil != err {
			panic(err)
		}
	})
	return db
}
