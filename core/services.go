package core

import (
	"github.com/NicolasMartino/simplebank/core/accounts"
	"github.com/NicolasMartino/simplebank/core/transfer"
	"github.com/NicolasMartino/simplebank/core/user"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type Services struct {
	AccountPersister  *accounts.AccountPersister
	AccountRetriever  *accounts.AccountRetriever
	TransferPersister *transfer.TransferPersister
	UserRetriever     *user.UserRetriever
	UserPersister     *user.UserPersister
}

//Services ctor with DI
func NewServices(store db.Store) *Services {
	accountPersister := accounts.NewAccountPersister(store)
	accountRetriever := accounts.NewAccountRetriever(store)
	transferPersister := transfer.NewTransferPersister(store, accountRetriever)
	userPstr := user.NewUserPersister(store)
	userRtvr := user.NewUserRetriever(store)

	services := &Services{
		AccountPersister:  accountPersister,
		AccountRetriever:  accountRetriever,
		TransferPersister: transferPersister,
		UserRetriever:     userRtvr,
		UserPersister:     userPstr,
	}
	return services
}
