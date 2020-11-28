// +build wireinject

package orm

import (
	"github.com/google/wire"
)

func InitialiseDb() Database {
	wire.Build(InitialiseConfig, providerGorm)
	return Database{}
}

func InitialiseConfig() config {
	wire.Build(wire.Value(config{
		Url:      "127.0.0.1",
		Port:     5432,
		Dbname:   "quiz",
		Ssl:      "disable",
		User:     "nishantagarwal",
		Password: " ",
	}))
	return config{}
}
