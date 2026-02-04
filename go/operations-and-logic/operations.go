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
		case 1:
			Temperature()
		case 2:
			RainingUmbrella()
		case 3:
			MemberDiscount()
		case 4:
			NumberComparison()
		case 5:
			NestedPuzzle()
		case 0:
			fmt.Println("Goodbye!")
			ifContinue = false
		}

	}
}

// Ask for a temperature, if it is over 80 print it is hot, if under 80 print it isn't hot
func Temperature() {
	fmt.Println("\nIs it hot or cold???")
	temperature := GetValidatedNumber("Enter a temperature: ", -50, 150)

	if temperature >= 80 {
		fmt.Println("It is hot!")
		fmt.Println("Temperature: ", temperature)
	} else {
		fmt.Println("It is not that hot...")
		fmt.Println("Temperature: ", temperature)
	}
}

// Ask user if it is raining or if you have an umbrella.
// Print if the user will stay dry according to their choices
func RainingUmbrella() {
	fmt.Println("\nRaining and Umbrella")

	fmt.Println("Is it raining?\n1 - Yes\n2 - No")
	userRaining := GetValidatedNumber("Enter option: ", 1, 2)

	fmt.Println("Do you have an umbrella?\n1 - Yes\n2 - No")
	userUmbrella := GetValidatedNumber("Enter option: ", 1, 2)

	if userRaining == 1 && userUmbrella == 2 {
		fmt.Println("You will get wet!")
	} else {
		fmt.Println("You will stay dry")
	}

}

// If a customer is a member -> 10% discount
// If they spend over $100 and are a member -> add 5% discount
// If they are not a member, they get no discount.
func MemberDiscount() {
	fmt.Println("\nMEMBER DISCOUNT")

	var isMember bool
	discount := 0.0

	userMember := GetValidatedNumber("Are you a member?\n1 - Yes\n2 - No\nEnter Option: ", 1, 2)
	switch userMember {
	case 1:
		isMember = true
	case 2:
		isMember = false
	}

	userBill := GetValidatedNumber("How much did you spend? (Round to the nearest dollar): $", 0, 10000)

	if isMember {
		discount += .10

		if userBill > 100 {
			discount += .05
		}
	}

	totalWithDiscount := float64(userBill) - (discount * float64(userBill))
	fmt.Println(strings.Repeat("-", 55))
	fmt.Printf("TOTAL WITH DISCOUNT: $%.2f\n", totalWithDiscount)
	fmt.Println(strings.Repeat("-", 55))
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
