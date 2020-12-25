//+build wireinject

package crm

import (
	"github.com/google/wire"

	"github.com/razorpay/quiz/package/orm"
)

func InitialiseServer() (Server, error) {
	wire.Build(InitialiseCore, provideServer)
	return Server{}, nil
}

func InitialiseCore() (Core, error) {
	wire.Build(InitlaiseRepo, provideCore)
	return Core{}, nil
}

func InitlaiseRepo() (Repo, error) {
	wire.Build(orm.InitialiseDb, provideRepo)
	return Repo{}, nil
}
