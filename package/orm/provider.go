package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	PostgresConnectionDSNFormat = "host=%s port=%d dbname=%s sslmode=%s user=%s password=%s"
)

func providerGorm(conf config) Database {
	db := Database{}
	path := fmt.Sprintf(PostgresConnectionDSNFormat, conf.Url, conf.Port, conf.Dbname, conf.Ssl, conf.User, conf.Password)
	dbInstance, err := gorm.Open("postgres", path)
	if err != nil {
		panic(err)
	}
	db.Db = dbInstance
	return db
}
