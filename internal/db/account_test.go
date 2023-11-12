package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/machingclee/2023-11-04-go-gin/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    sql.NullString{String: user.Username, Valid: true},
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
