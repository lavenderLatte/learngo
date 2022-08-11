// not compiling account.go, we will compile main.go

package accounts

import (
	"errors"
	"fmt"
)

// PUBLIC STRUCT: if struct name start with uppercase --> security issue!
type Account struct {
	// ownner  string // private data member
	// balance float32
	Owner   string // public data member
	Balance float32
}

var errLowBalance = errors.New("balance is lower than withdraw amount")

// PRIVATE STRUCT: if struct name starts with lowercase
type account struct {
	owner   string
	balance float32
}

// CONSTRUCTOR
// Make a function that returns mem addr of account obj instead of constructor
// NewAccount function takes owner, creates new account, and returns the newly created account object (not address)
// using * and & in order to not create a copy of object (slow)
func NewAccount(owner string) *account { // what will happen if I just return account? without *? func'll return copy of thisAccount
	thisAccount := account{owner: owner, balance: 0}
	return &thisAccount // &thisAccount returns addr/pointer to thisAccount == *account!!!! 엄청 헷갈리네 근데 중요해!
}

// METHODS
// Deposit is a public method
// This function won't add amount to balance!
// func (theAccount account) Deposit(amount float32) { // theAccount is type account and it's called receiver
// 	theAccount.balance += amount // ineffective assignment!: the amount won't added up to the balance!!
// 	// when struct object is passed to a method or function, Go makes copy of the object; theAccount is a copy of raviAccount!!
// }

// to solve problem above
func (theAccount *account) Deposit(amount float32) { // *account indicates that don't make copy, give me orignial receiver/object!!
	theAccount.balance += amount
}

func (theAccount account) ViewBalance() float32 {
	return theAccount.balance
}

func (theAccount *account) Withdraw(amount float32) error {
	// there is no error handling such as try-exception/catch in Go; have to manually check errors
	if amount > theAccount.balance {
		// return errors.New("can't go minus")
		return errLowBalance // can create error var and return that instead; better practice
	}
	theAccount.balance -= amount
	return nil // nil == null or none
}

func (theAccount *account) ChangeOwner(newOwner string) {
	theAccount.owner = newOwner
}

func (theAccount *account) ViewOwner() string {
	return theAccount.owner
}

// toString() method of Go
func (theAccount *account) String() string {
	return fmt.Sprint(theAccount.ViewOwner(), " has $", theAccount.ViewBalance(), " in her/his account.")
	// return fmt.Sprint(theAccount.owner, " has $", theAccount.balance, " in her/his account.")
}
