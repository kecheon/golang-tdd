package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit Bitcoin and get balance", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Stringify Bitcoin to print", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{}
		wallet.Deposit(10)
		b := wallet.Balance()
		got := b.String()
		want := "10 BTC"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw Bitcoin", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(50))
		assertNoError(t, err)
		got := wallet.Balance()
		want := Bitcoin(50)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(150))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet.Balance(), Bitcoin(100))
		if err == nil {
			t.Error("Error should occur but didn't")
		}
	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t *testing.T, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
