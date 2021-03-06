package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNotError(t, err)
	})

	t.Run("Withraw Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsuficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNotError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got and error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted and error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
