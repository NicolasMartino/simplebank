// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	FindAccount(ctx context.Context, id int64) (Account, error)
	FindAccountForUpdate(ctx context.Context, id int64) (Account, error)
	FindAccountsWithPagination(ctx context.Context, arg FindAccountsWithPaginationParams) ([]Account, error)
	FindEntriesWithPagination(ctx context.Context, arg FindEntriesWithPaginationParams) ([]Entry, error)
	FindEntry(ctx context.Context, id int64) (Entry, error)
	FindTransfer(ctx context.Context, id int64) (Transfer, error)
	FindUser(ctx context.Context, id int64) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateAccountAddToBalance(ctx context.Context, arg UpdateAccountAddToBalanceParams) (Account, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
	UpdateUserHash(ctx context.Context, arg UpdateUserHashParams) (User, error)
}

var _ Querier = (*Queries)(nil)
