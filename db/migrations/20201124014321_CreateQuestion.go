package migrations

import (
	"database/sql"

	"github.com/pressly/goose"

	crm "github.com/razorpay/quiz/crm"

	"github.com/razorpay/quiz/package/orm"
)

func init() {
	goose.AddMigration(Up_20201124014321, Down_20201124014321)
}

// Up is executed when this migration is applied
func Up_20201124014321(txn *sql.Tx) error {
	db := orm.InitialiseDb()
	db.Db.AutoMigrate(&crm.Question{})
	db.Db.Close()
	return nil
}

// Down is executed when this migration is rolled back
func Down_20201124014321(txn *sql.Tx) error {
	db := orm.InitialiseDb()
	db.Db.DropTableIfExists(&crm.Question{})
	db.Db.Close()
	return nil
}
