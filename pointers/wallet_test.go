package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Litcoin(10))

	got := wallet.Balance()
	want := Litcoin(10)

	if want != got {
		t.Errorf("Expected %s but got %s", want, got)
	}
}
