/**
 * @Author: mariomang
 * @Date: 2024-01-18 22:55:20
 * @File: db/postgres.go
**/

package repo

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var once sync.Once
var pgdb *PostgresRepo

type PostgresRepo struct {
	db *gorm.DB

	expiration time.Duration
	keyPrefix  string
}

func InitPostgresRepo(addr, username, password, dbname string, port int) {

	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			addr, username, password, dbname, port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:      nil,
			PrepareStmt: true,
		})
		if err != nil {
			log.Fatalln(err)
		}

		pgdb = &PostgresRepo{
			db:         db,
			expiration: 0,
			keyPrefix:  "",
		}
	})
}

func GetPostgresRepo() *PostgresRepo {
	if pgdb == nil {
		panic(errors.New("postgres connection not initialized"))
	}
	return pgdb
}
