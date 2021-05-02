package crm

type ICore interface {
	AddQuestion(q Question)
}

type Core struct {
	Repo IRepo
}

func (c *Core) AddQuestion(q Question) {
	c.Repo.AddQuestion(q)
}
