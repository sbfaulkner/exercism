package account

import "sync"

const testVersion = 1

// Account represents a bank account
type Account struct {
	balance int64
	closed  bool
	mutex   sync.Mutex
}

// Open creates a new account with the specified opening balance.
func Open(openingBalance int64) *Account {
	if openingBalance < 0 {
		return nil
	}
	account := new(Account)
	account.balance = openingBalance
	return account
}

// Balance returns the current account balance.
func (account *Account) Balance() (int64, bool) {
	defer func() { account.mutex.Unlock() }()
	account.mutex.Lock()

	if account.closed {
		return 0, false
	}
	return account.balance, true
}

// Deposit deposits the specified amount into the account and returns the updated balance.
// Specifying a negative amount will withdraw funds instead.
func (account *Account) Deposit(amount int64) (int64, bool) {
	defer func() { account.mutex.Unlock() }()
	account.mutex.Lock()

	if account.closed {
		return 0, false
	}
	if amount < 0 && -amount > account.balance {
		return 0, false
	}
	account.balance += amount
	return account.balance, true
}

// Close closes the account and returns the closing balance.
func (account *Account) Close() (int64, bool) {
	defer func() { account.mutex.Unlock() }()
	account.mutex.Lock()

	if account.closed {
		return 0, false
	}
	account.closed = true
	return account.balance, true
}
