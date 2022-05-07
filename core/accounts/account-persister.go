package accounts

import (
	"context"

	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type AccountPersister struct {
	store db.Store
}

func NewAccountPersister(store db.Store) *AccountPersister {
	AccountPersister := &AccountPersister{
		store: store,
	}

	return AccountPersister
}

func (accountPersister *AccountPersister) CreateAccount(ctx context.Context, UserID int64, currency string) (db.Account, error) {
	args := db.CreateAccountParams{
		UserID:   UserID,
		Currency: currency,
		Balance:  0,
	}
	return accountPersister.store.CreateAccount(ctx, args)
}

func (accountPersister *AccountPersister) UpdateAccount(ctx context.Context, ID int64, balance float64) (db.Account, error) {
	return accountPersister.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:      ID,
		Balance: balance,
	})
}

func (accountPersister *AccountPersister) DeleteAccount(ctx context.Context, ID int64) error {
	return accountPersister.store.DeleteAccount(ctx, ID)
}
