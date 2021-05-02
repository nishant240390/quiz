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
	path := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		conf.User,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.Dbname)
	fmt.Println("connection attempt")
	dbInstance, err := gorm.Open("postgres", path)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection successful")
	}
	db.Db = dbInstance
	return db
}
