package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type event struct {
	id            int
	beforeBalance int
	afterBalance  int
	operation     string
	timeStamp     int64
}

var ledger = make(map[string]*Account)

type Account struct {
	owner    string
	balance  int
	accNo    string
	passbook map[int]event
}

func (a *Account) Deposit(amount string) (*Account, error) {
	amt, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Invalid amount")
		return nil, err
	}
	beforeBalance := a.balance
	a.balance += amt
	fmt.Printf("Deposited %d to the account %q\n with current balance %d", amt, a.owner, a.balance)
	updatePassbook(*a, event{beforeBalance: beforeBalance, afterBalance: a.balance, operation: "deposit", timeStamp: time.Now().Unix()})
	return a, nil
}

func (a *Account) Withdraw(amount string) (*Account, error) {
	amt, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Invalid amount")
		return nil, err
	}
	if a.balance < amt {
		return nil, errorInsufficientFunds
	}
	beforeBalance := a.balance
	a.balance -= amt
	updatePassbook(*a, event{beforeBalance: beforeBalance, afterBalance: a.balance, operation: "withdraw", timeStamp: time.Now().Unix()})
	fmt.Printf("Withdrawn %d from the account %q\n", amt, a.owner)
	return a, nil

}

func printBalance(a Account) {
	fmt.Printf("The balance for the account %q is %d\n", a.owner, a.balance)
}

var errorInsufficientFunds = errors.New("insufficient funds")

func updatePassbook(a Account, e event) {
	if a.passbook == nil {
		a.passbook = make(map[int]event)
	}
	//add the event to passbook
	a.passbook[len(a.passbook)+1] = e
	fmt.Printf("Passbook updated for account %q\n with data %q", a.owner, e.operation)
}

func createAccount(owner string, balance int) *Account {
	var account = &Account{owner: owner, balance: balance, accNo: strconv.Itoa(len(ledger) + 1)}
	ledger[account.accNo] = account
	fmt.Printf("Account created for %q with balance %d\n and account number %q\n",
		account.owner, account.balance, account.accNo)
	updatePassbook(*account, event{beforeBalance: 0, afterBalance: balance, operation: "Account created", timeStamp: time.Now().Unix()})
	return ledger[account.accNo]
}
