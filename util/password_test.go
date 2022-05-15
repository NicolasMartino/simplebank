package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateHashAndCheckPassword(t *testing.T) {
	password := "1234"
	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	err = CheckPassword(hashedPassword, "worng password")
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

// salt is random => two hashes of the same password should be different
func TestTwoHashOfTheSamePasswordAreDifferent(t *testing.T) {
	password := "1234"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, password, hashedPassword2)

	require.NotEqual(t, hashedPassword, hashedPassword2)
}
