//+build wireinject

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
		Url:      "full_db_postgres",
		Port:     5432,
		Dbname:   "quiz",
		Ssl:      "disable",
		User:     "postgres",
		Password: "123",
	}))
	return config{}
}
