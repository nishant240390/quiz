package crm

import (
	`context`
	`fmt`
	
	crmv1 `github.com/razorpay/quiz/quiz/crm/v1`
)

type Server struct {}

func (s *Server) AddQuestion(context context.Context,request *crmv1.AddQuestionRequest) (*crmv1.AddQuestionResponse, error) {
	fmt.Println(" incoming request", request)
	return &crmv1.AddQuestionResponse{Success:true}, nil
}