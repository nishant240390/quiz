package main

import (
	"fmt"
	`net/http`
	
	"github.com/razorpay/quiz/crm"
	"github.com/razorpay/quiz/package/orm"
	crmv1 `github.com/razorpay/quiz/quiz/crm/v1`
)

func main() {
	db := orm.Database{}
	db.InitGorm()
	var enitities []interface{}
	enitities = append(enitities, &crm.Question{})
	db.MigrateAll(enitities)
	fmt.Println("entities added in db ")
	
	server := &crm.Server{}
	twirpHandler := crmv1.NewCrmServiceServer(server, nil)
	http.ListenAndServe( ":8009", twirpHandler)
}
