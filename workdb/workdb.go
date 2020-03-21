package workdb

import (
	"fmt"
	"server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// InitDB this function initials data base
func InitDB() {
	config := config.GetConfig("./config/config.json")
	connectStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.HostDb, config.PortDb, config.UserDb, config.NameDb, config.PasswordDb)
	db, err = gorm.Open("postgres", connectStr)
	// db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
}

// GetDB this function returns data base link
func GetDB() *gorm.DB {
	return db
}
