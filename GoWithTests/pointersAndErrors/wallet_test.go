package pointersAndErrors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWallet(t *testing.T) {
	verifyBalance := func(t *testing.T, wallet *Wallet, expected Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expected {
			t.Errorf("Balance was incorrect, got: %s, want: %s.", result.String(), expected.String())
		}
	}
	verifyError := func(t *testing.T, result error, expected error) {
		t.Helper()
		if result == nil {
			t.Fatal("expected an error but none was happen")
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Error was incorrect, got: %s, want: %s.", result, expected)
		}
	}
	verifyAnExpectedError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("an expected error occurred")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		verifyBalance(t, &wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		verifyBalance(t, &wallet, Bitcoin(10))
		verifyAnExpectedError(t, err)
	})
	t.Run("Withdraw wrong balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		valueWithdraw := Bitcoin(100)
		err := wallet.Withdraw(valueWithdraw)

		verifyBalance(t, &wallet, initialBalance)
		msgErr := fmt.Errorf("you cannot withdraw from %s to %s", initialBalance.String(), valueWithdraw.String())
		verifyError(t, err, msgErr)
	})
}
