package main

import (
	"fmt"

	"example.com/gobank/file_ops"
)

func Withdraw(account_balance *float64) {
	fmt.Print("| Your Withdraw: ")
	var withdraw_amount float64
	fmt.Scan(&withdraw_amount)
	if withdraw_amount <= 0 {
		fmt.Println("| Invalid: Must be Greater Than Zero")
	} else {
		if withdraw_amount > *account_balance {
			fmt.Println("| Invalid: You cannot withdraw more than you have.")
		} else {
			*account_balance -= withdraw_amount
			file_ops.Write_Float_To_File(*account_balance, ACCOUNT_BALANCE_FILE)
		}
	}
}

func Deposit(account_balance *float64) {
	fmt.Print("| Your Deposit: ")
	var deposit_amount float64
	fmt.Scan(&deposit_amount)
	if deposit_amount <= 0 {
		fmt.Println("| Invalid: Must be Greater Than Zero")
	} else {
		*account_balance += deposit_amount
		file_ops.Write_Float_To_File(*account_balance, ACCOUNT_BALANCE_FILE)
	}
}
