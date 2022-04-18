package core

import (
	"github.com/NicolasMartino/simplebank/core/accounts"
	"github.com/NicolasMartino/simplebank/core/transfer"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type Services struct {
	AccountPersister  *accounts.AccountPersister
	AccountRetriever  *accounts.AccountRetriever
	TransferPersister *transfer.TransferPersister
}

//Services ctor with DI
func NewServices(store db.Store) *Services {
	accountPersister := accounts.NewAccountPersister(store)
	accountRetriever := accounts.NewAccountRetriever(store)
	transferPersister := transfer.NewTransferPersister(store, accountRetriever)

	services := &Services{
		AccountPersister:  accountPersister,
		AccountRetriever:  accountRetriever,
		TransferPersister: transferPersister,
	}
	return services
}
