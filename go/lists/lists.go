package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("---SELECT WHICH EXERCISE TO SEE---")
		fmt.Println("1 - Color Check\n2 - List Modification\n3 - Total, Avg, Min, Max\n4 - Vowel Filter\n0 - Exit")
		userChoice := GetValidatedNumber("Enter Option: ", 0, 4)

		switch userChoice {
		case 1:
			colorCheck()
		case 2:
			listMod()
		case 3:
			calculations()
		case 4:
			vowelSort()
		case 0:
			fmt.Println("Goodbye!")
			ifContinue = false
		}
	}
}


// Create list of colors
// Ask user to enter a color
// Print message if their color is or isn't in the list
func colorCheck() {}

// Create a list of 5 numbers
// If the third number (num[2]) is less than 50, change it to 50
// If the total of the list is greater than 100, change the last number to 0
// Swap the first and third number
// Print the list
func listMod() {}

// Create a list of 10 numbers
// Calculate the total, average, minimum and maximum
// Print calculations
func calculations() {}

// Create list of 7 words
// For each vowel, print the number of vowels in that word
func vowelSort() {
	
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
