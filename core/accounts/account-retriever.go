package accounts

import (
	"context"
	"fmt"

	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type AccountRetriever struct {
	store db.Store
}

func NewAccountRetriever(store db.Store) *AccountRetriever {
	return &AccountRetriever{
		store: store,
	}
}

func (accountRetriever *AccountRetriever) RetrieveOneAccount(ctx context.Context, ID int64) (db.Account, error) {
	return accountRetriever.store.FindAccount(ctx, ID)
}

func (accountRetriever *AccountRetriever) RetrieveAccountsWithPagination(ctx context.Context, pageNumber int32, pageSize int32) ([]db.Account, error) {
	args := db.FindAccountsWithPaginationParams{
		Offset: (pageNumber - 1) * pageSize,
		Limit:  pageSize,
	}
	return accountRetriever.store.FindAccountsWithPagination(ctx, args)
}

func (accountRetriever *AccountRetriever) ValidateAccount(ctx context.Context, ID int64, currency string) (err error) {
	account, err := accountRetriever.RetrieveOneAccount(ctx, ID)

	if account.Currency != currency {
		err = fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
	}
	return
}
