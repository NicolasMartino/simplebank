package db

import (
	"context"
	"testing"
	"time"

	"github.com/NicolasMartino/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomHash(),
		FirstName:      util.RandomName(),
		LastName:       util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.FirstName, user.FirstName)
	require.Equal(t, args.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, time.Since(user.PasswordChangeAt))
	return user
}

func TestCreateUser(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	createRandomUser(t)
}
func TestGetUserById(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	expectedUser := createRandomUser(t)

	account, err := testQueries.FindUser(context.Background(), expectedUser.ID)
	require.NoError(t, err)

	require.NotZero(t, account.ID)
	require.Equal(t, expectedUser.ID, account.ID)
	require.Equal(t, expectedUser.Email, account.Email)
	require.Equal(t, expectedUser.FirstName, account.FirstName)
	require.Equal(t, expectedUser.LastName, account.LastName)
	require.Equal(t, expectedUser.HashedPassword, account.HashedPassword)
	require.Equal(t, expectedUser.CreatedAt, account.CreatedAt)
	require.Equal(t, expectedUser.PasswordChangeAt, account.PasswordChangeAt)

}

func TestUpdateHash(t *testing.T) {
	teardown := setupDBTestSuite(t)
	defer teardown(t)

	user := createRandomUser(t)

	updateCmd := UpdateUserHashParams{
		ID:             user.ID,
		HashedPassword: util.RandomHash(),
	}

	updatedUser, err := testQueries.UpdateUserHash(context.Background(), updateCmd)
	require.NoError(t, err)

	require.NotZero(t, updatedUser.ID)
	require.Equal(t, user.ID, updatedUser.ID)
	require.Equal(t, user.Email, updatedUser.Email)
	require.Equal(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, user.LastName, updatedUser.LastName)
	require.Equal(t, updateCmd.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, user.CreatedAt, updatedUser.CreatedAt)
	require.True(t, updatedUser.PasswordChangeAt.After(user.PasswordChangeAt))
}

// func TestDeleteUser(t *testing.T) {
// 	teardown := setupDBTestSuite(t)
// 	defer teardown(t)

// 	accountToDelete := createRandomAccount(t)

// 	err := testQueries.DeleteAccount(context.Background(), accountToDelete.ID)
// 	require.NoError(t, err)

// 	findDeletedaccount, err := testQueries.FindAccount(context.Background(), accountToDelete.ID)

// 	require.Empty(t, findDeletedaccount)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// }

// func TestListUsers(t *testing.T) {
// 	teardown := setupDBTestSuite(t)
// 	defer teardown(t)

// 	for i := 0; i < 10; i++ {
// 		createRandomAccount(t)
// 	}

// 	args := FindAccountsWithPaginationParams{
// 		Offset: 5,
// 		Limit:  5,
// 	}

// 	accounts, err := testQueries.FindAccountsWithPagination(context.Background(), args)
// 	require.NoError(t, err)
// 	require.Len(t, accounts, 5)

// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 	}
// }
