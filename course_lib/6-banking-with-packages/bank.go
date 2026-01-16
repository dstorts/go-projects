package main

import (
	"fmt"

	"example.com/gobank/file_ops"
	//above: import our custom package in this project.
	//       must start with the module name defined in go.mod
	//       must have the same name as the package name which matches its containing folder name
)

const ACCOUNT_BALANCE_FILE string = "Balance.txt"

func main() {
	var account_balance, err = file_ops.Get_Float_From_File(ACCOUNT_BALANCE_FILE)
	if err != nil {
		fmt.Println("Error reading Balance File: ")
		fmt.Println(err)
	}
	var choice int
	for { //this is how you define an infinite for loop
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
