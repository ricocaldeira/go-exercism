// Package account implements functions to deal with bank accounts
package account

import (
	"sync"
)

type Account struct {
	balance  int64
	isClosed bool
	mutex    sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, isClosed: false}
}
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	if a.isClosed {
		a.mutex.Unlock()
		return 0, false
	}
	a.isClosed = true
	payout, a.balance = a.balance, 0
	a.mutex.Unlock()
	return payout, true

}
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	if a.isClosed {
		a.mutex.Unlock()
		return 0, false
	}
	balance = a.balance
	a.mutex.Unlock()
	return balance, true
}
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	if a.isClosed {
		a.mutex.Unlock()
		return 0, false
	}
	newBalance = a.balance + amount
	if newBalance < 0 {
		a.mutex.Unlock()
		return 0, false
	}
	a.balance += amount
	a.mutex.Unlock()
	return newBalance, true
}
