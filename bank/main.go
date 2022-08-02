package main

import (
	"fmt"

	"github.com/panseung/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("panseung")
	account.Deposit(10)
	fmt.Println(account)
	
	err := account.Withdraw(7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account.Balance())
	}
}
