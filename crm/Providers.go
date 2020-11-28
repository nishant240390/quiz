package crm

import (
	"github.com/razorpay/quiz/package/orm"
)

func provideServer(core Core) Server {
	return Server{Core: &core}
}

func provideCore(repo Repo) Core {
	return Core{Repo: &repo}
}

func provideRepo(db orm.Database) Repo {
	return Repo{Db: db.Db}
}
