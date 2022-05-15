package db

import (
	"context"
	"testing"
	"time"

	"github.com/NicolasMartino/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	args := CreateUserParams{
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
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

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	user := createRandomUser(t)

	updateCmd := UpdateUserHashParams{
		ID:             user.ID,
		HashedPassword: hashedPassword,
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

// TODO add delete user with account test
