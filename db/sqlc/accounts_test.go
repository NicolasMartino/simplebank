package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/NicolasMartino/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	createRandomAccount(t)
}
func TestGetAccountById(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	expectedAccount := createRandomAccount(t)

	account, err := testQueries.FindAccount(context.Background(), expectedAccount.ID)
	require.NoError(t, err)

	require.Equal(t, expectedAccount.ID, account.ID)
	require.Equal(t, expectedAccount.Owner, account.Owner)
	require.Equal(t, expectedAccount.Balance, account.Balance)
	require.Equal(t, expectedAccount.Currency, account.Currency)
	require.Equal(t, expectedAccount.CreatedAt, account.CreatedAt)

	require.NotZero(t, account.ID)
}

func TestUpdateAccount(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	accountToUpdate := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      accountToUpdate.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, accountToUpdate.ID, account.ID)
	require.Equal(t, accountToUpdate.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, accountToUpdate.Currency, account.Currency)
	require.Equal(t, accountToUpdate.CreatedAt, account.CreatedAt)

	require.NotZero(t, account.ID)
}

func TestDeleteAccount(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	accountToDelete := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), accountToDelete.ID)
	require.NoError(t, err)

	findDeletedaccount, err := testQueries.FindAccount(context.Background(), accountToDelete.ID)

	require.Empty(t, findDeletedaccount)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccounts(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := FindAccountsWithPaginationParams{
		Offset: 5,
		Limit:  5,
	}

	accounts, err := testQueries.FindAccountsWithPagination(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
