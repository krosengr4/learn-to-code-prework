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
func colorCheck() {

	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("\t---COLOR CHECKER---")
	fmt.Println(strings.Repeat("-", 55))

	colorList := []string{"blue", "purple", "red", "magenta"}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a color: ")
	scanner.Scan()
	lowerColor := strings.ToLower(scanner.Text())
	userColor := strings.TrimSpace(lowerColor)

	isColorListed := false

	for _, color := range colorList {
		if color == userColor {
			isColorListed = true
		}
	}

	fmt.Println(strings.Repeat("-", 55))
	if isColorListed {
		fmt.Println("That color is in the list!!!")
	} else {
		fmt.Println("That color is not listed :(")
	}
	fmt.Println(strings.Repeat("-", 55))

	PauseScreen()
}

// Create a list of 5 numbers
// If the third number (num[2]) is less than 50, change it to 50
// If the total of the list is greater than 100, change the last number to 0
// Swap the first and third number
// Print the list
func listMod() {
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("\t---LIST MODIFICATION---")
	fmt.Println(strings.Repeat("-", 55))

}

// Create a list of 10 numbers
// Calculate the total, average, minimum and maximum
// Print calculations
func calculations() {
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("\t---LIST CALCULATIONS---")
	fmt.Println(strings.Repeat("-", 55))

	nums := [10]int{23, 64, 28, 33, 56, 97, 23, 18, 4, 67}
	total := 0
	var average float32
	min := 100000
	max := 0

	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("ORIGINAL LIST:", nums)
	fmt.Println(strings.Repeat("-", 55))

	for _, num := range nums {
		total += num

		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	
	average = float32(total) / float32(len(nums))

	fmt.Println(strings.Repeat("-", 55))
	fmt.Printf("TOTAL: %d\n", total)
	fmt.Printf("AVERAGE: %.2f\n", average)
	fmt.Printf("SMALLEST: %d\n", min)
	fmt.Printf("LARGEST: %d\n", max)
	fmt.Println(strings.Repeat("-", 55))

	PauseScreen()
}

// Create list of 7 words
// For each vowel, print the number of vowels in that word
func vowelSort() {
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("\t---VOWEL SORTER---")
	fmt.Println(strings.Repeat("-", 55))

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

// Helper function so that information pauses till user hits a key
func PauseScreen() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hit any key to continue...")
	scanner.Scan()
}
