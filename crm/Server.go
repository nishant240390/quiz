package crm

import (
	`context`
	
	crmv1 `github.com/razorpay/quiz/quiz/crm/v1`
)

type Server struct {
	Core ICore
}

func (s *Server) AddQuestion(context context.Context,request *crmv1.AddQuestionRequest) (*crmv1.AddQuestionResponse, error) {
	
	s.Core.AddQuestion(Question{
		Statement: request.Statement,
		Option1:   request.Option1,
		Option2:   request.Option2,
		Option3:   request.Option3,
		Option4:   request.Option4,
		Ans:       int(request.Answer),
	})
	
	return &crmv1.AddQuestionResponse{Success:true}, nil
}