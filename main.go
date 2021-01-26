package main

import (
	"banking/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("cham")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
