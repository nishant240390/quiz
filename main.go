package main

import (
	"net/http"

	"github.com/razorpay/quiz/crm"
	"github.com/razorpay/quiz/package/orm"
	crmv1 "github.com/razorpay/quiz/quiz/crm/v1"
)

func main() {
	db := orm.Database{}
	db.InitGorm()
	repo := crm.Repo{Db: db.Db}
	core := crm.Core{Repo: &repo}
	server := &crm.Server{Core: &core}
	twirpHandler := crmv1.NewCrmServiceServer(server, nil)
	http.ListenAndServe(":8009", twirpHandler)
}
