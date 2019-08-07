package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"sync"
)

type DB struct {
	Db *gorm.DB
}

var 	onceDbPostg sync.Once
var 	instanceDB *DB

// GetInstanceDb this connection for database
func GetInstanceDb() *gorm.DB {
	onceDbPostg.Do(func() {
		dbConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Config.Database.Host, Config.Database.Port, Config.Database.User, Config.Database.Password, Config.Database.Database)
		dbConnection, e := gorm.Open("postgres", dbConfig)
		if e == nil {
			fmt.Println("Connected to " + dbConfig)
			instanceDB = &DB{Db: dbConnection}
			dbConnection.LogMode(true)
		}else {
			panic(e)
		}

	})
	return instanceDB.Db
}

