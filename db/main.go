package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var Gormdb *gorm.DB

func init() {
	once.Do(func() {
		sqldb, err := sql.Open("pgx", "postgres://postgres:8249postgres..@47.93.142.176:5432/monitor")
		if err != nil {
			fmt.Println(err.Error())
		}

		Gormdb, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqldb,
		}), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error())
		}
	})
}
