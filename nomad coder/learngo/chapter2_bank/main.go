package main

import (
	"fmt"
	"learngo/learngo/chapter2_bank/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Onwer())
	fmt.Println(account)
}
