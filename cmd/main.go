package main

import (
	"fmt"

	"github.com/pressly/goose"

	"github.com/razorpay/quiz/package/orm"

	_ "github.com/razorpay/quiz/db/migrations"
)

var (
	dir = "db/migrations"
)

func main() {
	db := orm.InitialiseDb()
	arguments := []string{}

	err := goose.Run("up", db.Db.DB(), dir, arguments...)
	if err != nil {
		fmt.Print("migration could not complete", err)
		return
	}
}
