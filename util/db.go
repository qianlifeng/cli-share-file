package util

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path"
	"sync"
)

var (
	dbOnce       sync.Once
	dbConnection *gorm.DB
)

func getDBPath() string {
	return path.Join(GetAppFolder(), "tshare.db")
}

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		conn, err := gorm.Open(sqlite.Open(getDBPath()+"?_journal_mode=WAL&_synchronous=OFF"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(fmt.Sprintf("Can't connect to db , path=%s, err=%s ", getDBPath(), err.Error()))
		}
		dbConnection = conn
	})

	return dbConnection
}
