package db

import (
	"context"
	"testing"

	"github.com/NicolasMartino/simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestCreateTransferTx(t *testing.T) {
	//GIVEN
	teardown := setupDBTestSuite(t)

	defer teardown(t)

	store := NewStore(testDB)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	//run n concurrent transfer tx
	n := 25
	errsChannel := make(chan error)
	resultsChannel := make(chan TransferTxResult)
	amountChannel := make(chan float64)
	var accumulatedAmount float64
	accumulatedAmount = 0

	//WHEN async multiple transaction
	for i := 0; i < n; i++ {
		go func() {
			amount := util.RandomMoney()
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})
			errsChannel <- err
			resultsChannel <- result
			amountChannel <- amount
		}()
	}

	//THEN get channel message to check for errors and results
	for i := 0; i < n; i++ {
		err := <-errsChannel
		require.NoError(t, err)

		result := <-resultsChannel
		require.NotEmpty(t, result)

		amount := <-amountChannel
		require.NotZero(t, amount)
		accumulatedAmount = util.RoundToTwoDigits(accumulatedAmount + amount)

		transfer := result.Transfer
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		//check against db
		_, err = testQueries.FindTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = testQueries.FindEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toAccount.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = testQueries.FindEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//Check accounts
		fromAccountTx := result.FromAccount
		require.NotEmpty(t, fromAccountTx)

		toAccountTx := result.ToAccount
		require.NotEmpty(t, toAccountTx)

		//Check balance
		fromAccountBalanceDiff := util.RoundToTwoDigits(fromAccount.Balance - fromAccountTx.Balance)
		toAccountBalanceDiff := util.RoundToTwoDigits(toAccountTx.Balance - toAccount.Balance)

		require.Equal(t, fromAccountBalanceDiff, toAccountBalanceDiff)
		require.True(t, fromAccountBalanceDiff > 0)
		require.True(t, toAccountBalanceDiff > 0)
		require.Equal(t, accumulatedAmount, toAccountBalanceDiff)
		require.Equal(t, accumulatedAmount, fromAccountBalanceDiff)
	}

	// THEN check state after all asaync tx
	var updatedFromAccount Account
	var err error
	updatedFromAccount, err = testQueries.FindAccount(context.Background(), fromAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedFromAccount)

	var updatedToAccount Account
	updatedToAccount, err = testQueries.FindAccount(context.Background(), toAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedToAccount)

	expectedFromAccountBalance := util.RoundToTwoDigits(fromAccount.Balance - accumulatedAmount)
	expectedToAccountBalance := util.RoundToTwoDigits(toAccount.Balance + accumulatedAmount)
	require.Equal(t, expectedFromAccountBalance, util.RoundToTwoDigits(updatedFromAccount.Balance))
	require.Equal(t, expectedToAccountBalance, util.RoundToTwoDigits(updatedToAccount.Balance))
}

func TestTransferTxDeadlock(t *testing.T) {
	//GIVEN
	teardown := setupDBTestSuite(t)

	defer teardown(t)

	store := NewStore(testDB)

	accountA := createRandomAccount(t)
	accountB := createRandomAccount(t)

	//run n concurrent transfer tx
	n := 20
	errsChannel := make(chan error)
	amount := util.RandomMoney()

	//WHEN async multiple transaction
	for i := 0; i < n; i++ {
		var fromAccount Account
		var toAccount Account

		if i%2 == 1 {
			fromAccount = accountA
			toAccount = accountB
		} else {
			fromAccount = accountB
			toAccount = accountA
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})
			errsChannel <- err
		}()
	}

	//THEN get channel message to check for errors and results
	for i := 0; i < n; i++ {
		err := <-errsChannel
		require.NoError(t, err)
	}

	// THEN check state after all asaync tx
	var updatedAccountA Account
	var err error
	updatedAccountA, err = testQueries.FindAccount(context.Background(), accountA.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccountA)

	var updatedAccountB Account
	updatedAccountB, err = testQueries.FindAccount(context.Background(), accountB.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccountB)

	require.Equal(t, accountA.Balance, util.RoundToTwoDigits(updatedAccountA.Balance))
	require.Equal(t, accountB.Balance, util.RoundToTwoDigits(updatedAccountB.Balance))
}
