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
		fmt.Println("1 - Temperature\n2 - Raining or Umbrella\n3 - Store Discount\n4 - Min and Max Number\n5 - Nested Logic Puzzle\n0 - Exit")
		userChoice := GetValidatedNumber("Enter option: ", 0, 5)

		switch userChoice {

		}

	}
}

// Ask for a temperature, if it is over 80 print it is hot, if under 80 print it isn't hot
func Temperature() {
	fmt.Println("\nIs it hot or cold???")

}

// Ask user if it is raining or if you have an umbrella.
// Print if the user will stay dry according to their choices
func RainingUmbrella() {

}

// If a customer is a member -> 10% discount
// If they spend over $100 -> add 5% discount
func MemberDiscount() {

}

// Ask user for 3 numbers, determine the largest and smallest without built in functions
func NumberComparison() {

}

// Ask for four integer variables
// Print the second to largest number and if it is even or odd
func NestedPuzzle() {

}

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
