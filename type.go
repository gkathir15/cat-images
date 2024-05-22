package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Balance   int64  `json:"balance"`
	Number    int64  `json:"number"`
}

func NewAccount(firstname string, lastname string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		Firstname: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(1000000)),
	}
}
