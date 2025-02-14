package pointers

import "testing"

func assertError(t testing.TB, realErr error, expectedErr error) {
	t.Helper()
	if realErr == nil {
		t.Fatal("Expected an error but was given nil")
	}

	if realErr != expectedErr {
		t.Errorf("Expected error with reason %s, but recieved %s", realErr, expectedErr)
	}
}

func assertBalance(t testing.TB, wallet Wallet, expectedBalance Litcoin) {
	t.Helper()
	if wallet.balance != expectedBalance {
		t.Errorf("Expected %s but got %s", expectedBalance, wallet.balance)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Litcoin(10))

		assertBalance(t, wallet, Litcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Litcoin(8)}
		wallet.Withdraw(Litcoin(3))

		assertBalance(t, wallet, Litcoin(5))
	})

	t.Run("withdraw errors with insufficient balance", func(t *testing.T) {
		startingBalance := Litcoin(7)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Litcoin(9))

		assertError(t, err, ErrInsufficentFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
