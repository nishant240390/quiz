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
		Url:      "192.168.0.2",
		Port:     5432,
		Dbname:   "fullstack_api",
		Ssl:      "disable",
		User:     "steven",
		Password: "password",
	}))
	return config{}
}
