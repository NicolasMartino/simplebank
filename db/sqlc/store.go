package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplt_bank/util"
)

// This Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db, Queries: New(db)}
}

//Executes a function in the context of a DB transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rollbackError)
		}
		return err
	}

	return tx.Commit()
}

// Contain the input parameters of a trasfer transaction
type TransferTxParams struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

// Contains the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
}

// Transfer executes a money transfer between two accounts.
// It create a transfer record, add acount entries and update account balance within a signle DB transaction
func (store *Store) TransferTx(context context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(context, func(q *Queries) error {
		var err error
		roundedAmount := util.RoundToTwoDigits(arg.Amount)

		result.Transfer, err = q.CreateTransfer(context, CreateTransferParams(arg))

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(context, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -roundedAmount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(context, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    roundedAmount,
		})

		if err != nil {
			return err
		}

		result.FromAccount, result.ToAccount, err = makeSafeAccountUpdate(context, q, arg.FromAccountID, arg.ToAccountID, roundedAmount)

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

// Deadlock safe account update
func makeSafeAccountUpdate(context context.Context,
	q *Queries,
	fromAccountId int64,
	toAccountId int64,
	amount float64) (updatedFromAccount Account, updatedToAccount Account, err error) {
	if fromAccountId < toAccountId {
		updatedFromAccount, updatedToAccount, err = accountsUpdate(context, q, fromAccountId, -amount, toAccountId, amount)
	} else {
		updatedToAccount, updatedToAccount, err = accountsUpdate(context, q, toAccountId, amount, fromAccountId, -amount)
	}
	return
}

func accountsUpdate(
	context context.Context,
	q *Queries,
	firstAccountId int64,
	firstAmount float64,
	secondAccountId int64,
	secondAmount float64) (updatedFirstAccount Account, updatedSecondAccount Account, err error) {

	var firstAccount Account
	firstAccount, err = q.FindAccountForUpdate(context, firstAccountId)

	if err != nil {
		return
	}

	updatedFirstAccount, err = q.UpdateAccount(context, UpdateAccountParams{
		ID:      firstAccountId,
		Balance: firstAccount.Balance + firstAmount,
	})

	if err != nil {
		return
	}

	var secondAccount Account
	secondAccount, err = q.FindAccountForUpdate(context, secondAccountId)
	if err != nil {
		return
	}

	updatedSecondAccount, err = q.UpdateAccount(context, UpdateAccountParams{
		ID:      secondAccount.ID,
		Balance: secondAccount.Balance + secondAmount,
	})

	if err != nil {
		return
	}

	return
}
