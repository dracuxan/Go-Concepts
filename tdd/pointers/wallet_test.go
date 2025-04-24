package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	checkBitcoin := func(t testing.TB, w *Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	checkError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if err != want {
			t.Errorf("got err: %q, want %q", err, want)
		}
	}

	checkNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatalf("Ecountered an unexpected error %q", got)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		checkBitcoin(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(22))
		want := Bitcoin(78)

		checkNoError(t, err)
		checkBitcoin(t, &wallet, want)
	})

	t.Run("withdraw exceeding limits", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkError(t, err, ErrInsufficientFunds)
		checkBitcoin(t, &wallet, startingBalance)
	})
}
