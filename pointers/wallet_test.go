package pointers

import "testing"

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error but was given nil")
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

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
