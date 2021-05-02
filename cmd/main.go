package main

import (
	"fmt"
	"os"

	"github.com/pressly/goose"

	"github.com/razorpay/quiz/package/orm"

	_ "github.com/razorpay/quiz/db/migrations"
)

var (
	dir = "../db/migrations"
)

func main() {
	fmt.Print("\n init migration \n")
	db := orm.InitialiseDb()
	arguments := []string{}
	fmt.Print("\n running migration \n")

	if len(os.Args) < 2 {
		panic("no valid command")
	}
	err := goose.Run(os.Args[1], db.Db.DB(), dir, arguments...)
	if err != nil {
		fmt.Print("migration could not complete", err)
		return
	}
	fmt.Print("\n completed migration \n")
}
