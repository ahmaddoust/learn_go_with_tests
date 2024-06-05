package main

import "fmt"

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Withdraw subtracts some Bitcoin from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
