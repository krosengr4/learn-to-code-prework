package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("Select which excersise to see")
		fmt.Println("1 - Print numbers 1-10\n2 - Sum and average\n3 - FizzBuzz\n4 - Tax Calculation\n5 - Sorting list\n0 - Exit")
		userChoice := GetValidatedNumber("Enter option: ", 0, 5)

		switch userChoice {
		case 1:
			PrintNumbers()
		case 2:
			SumAndAvg()
		case 3:
			FizzBuzz()
		case 4:
			TaxCalculation()
		case 5:
			Sorting()
		case 0:
			fmt.Println("Goodbye!")
			ifContinue = false
		}
	}

}

// Loop through and print a list of  numbers
func PrintNumbers() {}

// Compute the sum and the average of a list of 5 numbers
func SumAndAvg() {}

// Given a list of 0 - 100
// If a number is divisible by 15 print fizzbuzz
// If a number is divisible by 5, print fizz
// If a number is divisible by 3, print buzz
func FizzBuzz() {}

// Given a list of 5 item prices
// For each price, compute the price with tax
// Print the total price along with each price with tax
//(Optional) Print and calculate the average cost with tax
func TaxCalculation() {}

// Given a list of 5 numbers, print the list in ascending order
func Sorting() {}

// Helper func to prompt user for integer type response, return that integer
func GetValidatedNumber(prompt string, min, max int) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if num, err := strconv.Atoi(input); err == nil {
			if num >= min && num <= max {
				return num
			}
			fmt.Printf("Number must be between %d and %d. Try again.\n", min, max)
		} else {
			fmt.Println("Invalid number. Please try again.")
		}
	}
}
