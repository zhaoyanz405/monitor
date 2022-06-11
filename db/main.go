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
		sqldb, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s/monitor", CurConfig.pgUser, CurConfig.pgPass, CurConfig.pgHost))
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
