package core

import (
	"github.com/NicolasMartino/simplebank/core/accounts"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type Services struct {
	AccountPersister *accounts.AccountPersister
	AccountRetriever *accounts.AccountRetriever
}

//Services ctor with DI
func NewServices(store db.Store) *Services {
	services := &Services{
		AccountPersister: accounts.NewAccountPersister(store),
		AccountRetriever: accounts.NewAccountRetriever(store),
	}
	return services
}
