package main

import (
	"net/http"

	"github.com/razorpay/quiz/crm"
	crmv1 "github.com/razorpay/quiz/quiz/crm/v1"
)

func main() {
	server, _ := crm.InitialiseServer()
	twirpHandler := crmv1.NewCrmServiceServer(&server, nil)
	http.ListenAndServe(":8009", twirpHandler)
}
