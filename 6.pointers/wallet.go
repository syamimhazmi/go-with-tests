package main

import (
	"errors"
	"fmt"
)

// In Go we can defined a struct with any type
type Bitcoin int

type Wallet struct {
	balance Bitcoin // then use the struct in another struct data type
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// instead use normal stuct type
// we add a pointer `*T` at the back of the type
// this will tell the code, that we want to manipulate the value of struct attribute
// the term that we are using is `struct pointers` or automatically dereferenced
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}

// when declaring a variable using `var` keyword,
// the variable is can be use globally across all packages
var ErrInsufficientFunds = errors.New("insufficient amount to withdraw")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

func main() {

}
