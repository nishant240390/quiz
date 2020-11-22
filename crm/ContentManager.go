package crm

type ICore interface {
	AddQuestion(q Question)
}

type Core struct {
	repo IRepo
}

func (c *Core) AddQuestion(q Question) {
	c.repo.AddQuestion()
}
