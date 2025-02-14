package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10.0

	if want != got {
		t.Errorf("Expected %.2f but got %.2f", want, got)
	}
}
