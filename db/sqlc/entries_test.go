package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    10,
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	return entry
}
func TestCreateEntry(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	entry := createRandomEntry(t)
	require.NotEmpty(t, entry)
}
func TestGetEntry(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	createdEntry := createRandomEntry(t)

	entry, err := testQueries.FindEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, createdEntry.AccountID, entry.AccountID)
	require.Equal(t, createdEntry.Amount, entry.Amount)
}
func TestDeleteEntry(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	createdEntry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)

	findDeletedEntry, err := testQueries.FindEntry(context.Background(), createdEntry.ID)
	require.Empty(t, findDeletedEntry)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestUpdateEntry(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	createdEntry := createRandomEntry(t)

	args := UpdateEntryParams{
		ID:     createdEntry.ID,
		Amount: 139,
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)
	require.Equal(t, createdEntry.ID, updatedEntry.ID)
	require.Equal(t, args.Amount, updatedEntry.Amount)
	require.Equal(t, createdEntry.CreatedAt, updatedEntry.CreatedAt)
	require.Equal(t, createdEntry.AccountID, updatedEntry.AccountID)
}
