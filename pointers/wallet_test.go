package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Litcoin(10))

	got := wallet.Balance()
	want := Litcoin(10)

	if want != got {
		t.Errorf("Expected %s but got %s", want, got)
	}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Litcoin(8)}
		wallet.Withdraw(Litcoin(3))

		got := wallet.Balance()
		want := Litcoin(5)

		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}
	})
}
