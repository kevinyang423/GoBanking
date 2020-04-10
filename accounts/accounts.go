package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type account struct {
	owner   string
	balance int
}

//NewAccount creates Account
func NewAccount(owner string) *account {
	ac := account{owner: owner, balance: 0}
	return &ac
}

//Deposit * amount on your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your Account
func (a account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("can't, not enough money")

//Withdraw from your Account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the Account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}
func (a account) Owner() string {
	return a.owner
}
func (a account) String() string {
	return fmt.Sprint(a.Owner(), "s account. \n Has", a.Balance())
}
