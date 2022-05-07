package accounts

import (
	"context"
	"fmt"
	"testing"

	mockDB "github.com/NicolasMartino/simplebank/db/mock"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
	"github.com/NicolasMartino/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name              string
	accountId         int64
	buildStubs        func(store *mockDB.MockStore)
	checkReturnValues func(t *testing.T, actualAccount db.Account, err error)
}

func TestFindOneAccountService(t *testing.T) {
	account := randomAccount()
	var accounts []db.Account
	for i := 0; i < 10; i++ {
		accounts = append(accounts, randomAccount())
	}

	testCases := []testCase{
		{
			name:      "account found",
			accountId: account.ID,
			buildStubs: func(store *mockDB.MockStore) {
				store.EXPECT().
					FindAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkReturnValues: func(t *testing.T, actualAccount db.Account, err error) {
				require.Nil(t, err)
				require.Equal(t, account.ID, actualAccount.ID)
				require.Equal(t, account.UserID, actualAccount.UserID)
				require.Equal(t, account.Balance, actualAccount.Balance)
				require.Equal(t, account.Currency, actualAccount.Currency)
			},
		},
		{
			name:      "account not found",
			accountId: account.ID,
			buildStubs: func(store *mockDB.MockStore) {
				store.EXPECT().
					FindAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkReturnValues: func(t *testing.T, actualAccount db.Account, err error) {
				require.Nil(t, err)
				require.Equal(t, account.ID, actualAccount.ID)
				require.Equal(t, account.UserID, actualAccount.UserID)
				require.Equal(t, account.Balance, actualAccount.Balance)
				require.Equal(t, account.Currency, actualAccount.Currency)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrler := gomock.NewController(t)
			defer ctrler.Finish()

			store := mockDB.NewMockStore(ctrler)
			tc.buildStubs(store)
			// Start test
			fmt.Printf("tescase n° %v", i)
			accountRetriever := NewAccountRetriever(store)
			actualAccount, err := accountRetriever.RetrieveOneAccount(context.Background(), tc.accountId)
			tc.checkReturnValues(t, actualAccount, err)

		})
	}

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		UserID:   util.RandomInt(1, 1000),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
