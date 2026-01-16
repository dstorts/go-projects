package main

import (
	"errors"
	"fmt"
	"os"
)

const ACCOUNT_BALANCE_FILE string = "Balance.txt"

func main() {
	var account_balance, err = Get_Balance_From_File()
	if err != nil {
		fmt.Println("Error reading Balance File: ")
		fmt.Println(err)
	}
	var choice int
	//for i := 0; i < 100; i++ { //this is how you would define a normal for loop with iterator and condition
	//for <CONDITION> { 		//this would loop until the given condition yeilds 'false'
	for { //this is how you define an infinite for loop
		Write_Balance_To_File(account_balance)

		Print_Options()

		fmt.Print("Input:")
		fmt.Scan(&choice)

		if choice == 7 {
			fmt.Println("LUCKY BOI!")
			// return
			continue // skips the rest of the current loop iteration and starts the next loop iteration. Special keyword for inside of for loops.
		}
		if choice == 666 {
			panic("The Devil has come for your soul!") // crashes the program with the given message
		}

		/*
			if choice == 1 {
				fmt.Println("Your balance is: ", account_balance)
			} else if choice == 2 {
				Deposit(&account_balance)
			} else if choice == 3 {
				Withdraw(&account_balance)
			} else {
				break //break out of the loop, not just this if else
			}
		*/
		switch choice {
		case 1:
			fmt.Println("| Your balance is: ", account_balance)
		case 2:
			Deposit(&account_balance)
		case 3:
			Withdraw(&account_balance)
		default:
			fmt.Println("Thank for Banking with Us today!")
			return
		}
	}
}

func Write_Balance_To_File(account_balance float64) {
	balanceText := fmt.Sprint(account_balance)
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balanceText), 0644)
	// above: []byte(balanceText) turns the string into a byte slice, which is what WriteFile needs
	// 0644 is the file permission setting, owner can read/write, group can read, others can read
}

func Get_Balance_From_File() (float64, error) {
	balanceText, err := os.ReadFile(ACCOUNT_BALANCE_FILE)
	// above: '_' is used to ignore the error value returned by ReadFile
	var balance float64
	if err != nil {
		//here if failed to read file
		return 1.0, errors.New("Failed to find Balance File.")
	} else {
		fmt.Sscan(string(balanceText), &balance)
	}
	return balance, nil
}

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
		Write_Balance_To_File(*account_balance)
	}
}

func Print_Options() {
	fmt.Println("|=========================|")
	fmt.Println("|---Welcome to Go Bank----|")
	fmt.Println("| What do you want to do? |")
	fmt.Println("| 1. Check Balance        |")
	fmt.Println("| 2. Deposity Money       |")
	fmt.Println("| 3. Withdraw Money       |")
	fmt.Println("| 4. Exit                 |")
	fmt.Println("|=========================|")
}
