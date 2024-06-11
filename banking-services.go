package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(bankingHandler)))
}

func bankingHandler(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/bank") {
		w.Write([]byte("Welcome to Go banking\n" +
			"Use /create/$name to create an account\n" +
			"Use /deposit/$accNo/$amount to deposit money\n" +
			"Use /withdraw/$accNo/$amount to withdraw money\n" +
			"Use /accounts to view all accounts\n" +
			"Use /ledger to view all accounts with balance\n" +
			"Use /balance/$accNo to view balance of an account\n" +
			"use /passbook/$accNo to view passbook of an account\n"))

		return
	}

	// account endpoint,returns all the accounts in ledger
	if strings.HasPrefix(r.URL.Path, "/accounts") {
		response := "Accounts in Ledger\n"
		for _, account := range ledger {
			response += fmt.Sprintf("Account Number: #%q, Owner: %s", account.accNo, account.owner)
		}
		w.Write([]byte(response))
		return
	}

	// Ledger endpoint prints all the accounts in ledger with balance
	if strings.HasPrefix(r.URL.Path, "/ledger") {
		response := "Accounts in Ledger\n"
		for _, account := range ledger {
			response += fmt.Sprintf("Account Number: %q, Owner: %s, Balance: %d\n", account.accNo, account.owner, account.balance)
		}
		w.Write([]byte(response))
		return
	}
	if strings.HasPrefix(r.URL.Path, "/create") {
		pathParams := strings.Split(r.URL.Path, "/")
		println("creating account for", pathParams[2])
		acc := createAccount(pathParams[2], 0)
		// print all account with details from ledger
		w.Write([]byte(fmt.Sprintf("Account created for %q with balance $%d\n and account number #%q\n",
			acc.owner, acc.balance, acc.accNo)))

		return
	}

	if strings.HasPrefix(r.URL.Path, "/deposit") {
		pathParams := strings.Split(r.URL.Path, "/")
		accNo := pathParams[2]
		amount := pathParams[3]
		println("depositing", amount, "to account", accNo)
		acc := ledger[accNo]
		account, err := acc.Deposit(amount)
		if err == nil {
			w.Write([]byte(fmt.Sprintf("Deposited $%s to the account %q\n", amount, account.owner)))
		} else {
			w.Write([]byte(fmt.Sprintf("Failed to deposit $%s to the account %q\n", amount, account.owner)))
		}
	}

	if strings.HasPrefix(r.URL.Path, "/withdraw") {
		pathParams := strings.Split(r.URL.Path, "/")
		accNo := pathParams[2]
		amount := pathParams[3]
		println("withdrawing", amount, "from account", accNo)
		acc := ledger[accNo]
		account, err := acc.Withdraw(amount)
		if err == nil {
			w.Write([]byte(fmt.Sprintf("Withdrawn $%s from the account %q\n", amount, account.owner)))
		} else {
			w.Write([]byte(fmt.Sprintf("Failed to withdraw $%s from the account %q\n", amount, account.owner)))
		}

	}

	if strings.HasPrefix(r.URL.Path, "/balance") {
		pathParams := strings.Split(r.URL.Path, "/")
		accNo := pathParams[2]
		println("checking balance for account", accNo)
		acc := ledger[accNo]
		printBalance(*acc)
		w.Write([]byte(fmt.Sprintf("The balance for the account %q is $%d\n", acc.owner, acc.balance)))

	}

	if strings.HasPrefix(r.URL.Path, "/passbook") {
		pathParams := strings.Split(r.URL.Path, "/")
		accNo := pathParams[2]
		println("checking passbook for account", accNo)
		acc := ledger[accNo]
		// print all the events of the passbook
		response := fmt.Sprintf("Passbook for account %q\n", acc.owner)
		for _, v := range acc.passbook {
			response += fmt.Sprintf("#%d at %s Operation: %s, Before: $%d, After: $%d\n", v.id,
				time.Unix(v.timeStamp, 0), v.operation, v.beforeBalance, v.afterBalance)
		}
		response += fmt.Sprintf("Current Balance: $%d\n", acc.balance)
		println("passbook data\n", response)
		w.Write([]byte(response))
		return

	}
}

func inProgress(w http.ResponseWriter) {
	w.Write([]byte("implementation in progress\n"))
}
