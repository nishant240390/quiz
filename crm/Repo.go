package crm

import (
	`github.com/jinzhu/gorm`
)

type IRepo interface {
	AddQuestion(q Question)
}
type Repo struct {
	Db *gorm.DB
}

func (r *Repo) AddQuestion(q Question) {
	r.Db.Create(&q)
}
