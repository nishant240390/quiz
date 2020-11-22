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

func (d *Database) InitGorm() {
	path := fmt.Sprintf(PostgresConnectionDSNFormat, "127.0.0.1", 5432, "quiz", "disable", "nishantagarwal", " ")
	db, err := gorm.Open("postgres", path)
	if err != nil {
		fmt.Println("error while opening db ", err)
	}
	d.Db = db
}

func (d *Database) MigrateAll(entities []interface{}) {
	for _, v := range entities {
		d.Db.DropTableIfExists(v)
		d.Db.AutoMigrate(v)
	}
}
