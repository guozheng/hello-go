package pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T)  {
	testBalance := func(t *testing.T, wallet Wallet, want int) {
		t.Helper()
		if wallet.Balance() != want {
			t.Errorf("got %d want %d", wallet.balance, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{0}
		w.Deposit(10)
		testBalance(t, w, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{10}
		err := w.Withdraw(5)
		assert.Nil(t, err, "should have enough balance to withdraw")
		testBalance(t, w, 5)
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		w := Wallet{10}
		err := w.Withdraw(50)
		assert.NotNil(t, err, "withdraw over balance should give error")
	})

}
