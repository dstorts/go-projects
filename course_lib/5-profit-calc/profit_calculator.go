package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, Rerr := getUserInput("Revenue: ")
	expenses, Eerr := getUserInput("Expenses: ")
	taxRate, Terr := getUserInput("Tax Rate: ")

	if Rerr != nil {
		fmt.Println("Error reading Revenue: ", Rerr)
		return
	}
	if Eerr != nil {
		fmt.Println("Error reading Expenses: ", Eerr)
		return
	}
	if Terr != nil {
		fmt.Println("Error reading Tax Rate: ", Terr)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt) //ebt in finance means Earnings Before Tax
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio) //ratio of EBT to Profit
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	//and of the Sprint functions return a string instead of printing a given string to console.
	//the 'f' in Sprintf allows formatting like in Printf.
	//%.3f means float with 3 decimal places.
	os.WriteFile("Results.txt", []byte(results), 0644)
	//[]byte is a 'slice' of bytes, it is like and array but the size is dynamic (grow/shrink as needed).
}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
