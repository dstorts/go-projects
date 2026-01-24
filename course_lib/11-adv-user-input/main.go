package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pure_input, carriage_input, scrubbed_input := Get_User_Input("Type in some multi-word input: ")

	fmt.Println("Pure Input ---")
	fmt.Print(pure_input)
	fmt.Print(pure_input)

	fmt.Println("Carriage Input ---")
	fmt.Print(carriage_input)
	fmt.Print(carriage_input)

	fmt.Println("Scrubbed Input ---")
	fmt.Print(scrubbed_input)
	fmt.Print(scrubbed_input)
}
func Get_User_Input(prompt string) (pure string, carriage string, scrubbed string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", "", ""
	}

	pure = input
	// Remove the newline character from the input using strings package
	progress := strings.TrimSuffix(input, "\n")
	// On Windows, also trim carriage return
	carriage = progress
	scrubbed = strings.TrimSuffix(progress, "\r")

	return input, carriage, scrubbed
}
