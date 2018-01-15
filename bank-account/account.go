// Package account implements functions to deal with bank accounts
package account

type Account struct {
	balance int64
	status  string
}

func (acc *Account) Close() (payout int64, ok bool) {
	acc.status = "closed"
	return acc.balance, !ok
}

func (acc *Account) Balance() (balance int64, ok bool) {
	ok = acc.status != "closed"
	if acc.balance < 0 {
		acc.balance = 0
	}

	return acc.balance, ok
}

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.balance += amount
	ok = acc.balance >= 0 && acc.status != "closed"
	return acc.balance, ok
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	acc := new(Account)
	acc.balance = initialDeposit
	acc.status = "opened"
	return acc
}
