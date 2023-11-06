package db

import (
	"context"
	"github.com/machingclee/2023-11-04-go-gin/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Owner, account.Owner)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
