package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	PostgresConnectionDSNFormat = "host=%s port=%d dbname=%s sslmode=%s user=%s password=%s"
)

type Database struct {
	Db *gorm.DB
}

func InitGorm() Database {
	db := Database{}
	path := fmt.Sprintf(PostgresConnectionDSNFormat, "127.0.0.1", 5432, "quiz", "disable", "nishantagarwal", " ")
	dbInstance, err := gorm.Open("postgres", path)
	if err != nil {
		fmt.Println("error while opening db ", err)
	}
	db.Db = dbInstance
	return db
}
