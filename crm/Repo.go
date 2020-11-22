package crm

import (
	"fmt"
)

type IRepo interface {
	AddQuestion()
}
type Repo struct {
}

func (r *Repo) AddQuestion() {
	fmt.Println("add questions")
}
