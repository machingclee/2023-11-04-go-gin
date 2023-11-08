package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errChan := make(chan error)
	resultChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errChan <- err
			resultChan <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)

		result := <-resultChan
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		resultFromAcocunt := result.FromAccount
		resultToAccount := result.ToAccount

		fmt.Println("before tx:", "fromAccount", account1.Balance, "toAccount", account2.Balance, "amount", int64(i+1)*amount)
		fmt.Println("after tx: ", "fromAccount", resultFromAcocunt.Balance, "toAccount", resultToAccount.Balance)
		fmt.Println("-------")

		diff1 := account1.Balance - resultFromAcocunt.Balance
		diff2 := resultToAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1 > 0)
	}
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 10
	amount := int64(10)

	errChan := make(chan error)
	// resultChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		var sign = int64(1)
		if i%2 == 0 {
			sign = int64(-1)
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        sign * amount,
			})
			errChan <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)
	}

	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
