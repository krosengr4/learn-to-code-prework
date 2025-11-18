package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Exercise 1: Algorithms

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("Enter number (1-3) to view which exercise you would like to see.")
		fmt.Println("1 - Vote Eligibility\n2 - Tip Computation\n3 - Layaway Computation\n0 - Exit")
		userChoice := GetValidatedNumber("Enter option: ", 0, 3)

		switch userChoice {
		case 1:
			age := GetValidatedNumber("Enter an age (1-100): ", 1, 100)
			voteEligibility(age)

		case 2:
			mealPrice := GetValidatedNumber("Enter a price for a resturaunt bill: ", 0, 10000)
			floatMealPrice := float64(mealPrice)
			tipComputation(floatMealPrice)

		case 3:
			itemPrice := GetValidatedNumber("Enter the price for an item: ", 0, 10000)
			months := GetValidatedNumber("Enter the number of months: ", 1, 500)

			floatItemPrice := float64(itemPrice)

			layawayComputation(floatItemPrice, months)
		case 0:
			fmt.Println("Goodbye!")
			ifContinue = false
		}
	}

}

// Vote Eligibility: Given an age, determine if the person if eligible to vote.
func voteEligibility(age int) {
	fmt.Println("Vote Eligibility. Age:", age)
	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else {
		fmt.Println("You are NOT eligible to vote :(")
	}
}

// Tip Computation: Given a bill for a meal at a resturaunt, ask the user for a tip percentage, and compute the price of the tip
func tipComputation(price float64) {

	tip := GetValidatedNumber("Enter the tip percent you want to give:", 0, 100)

	tipAsDecimal := float64(tip) / 100

	amountOfTip := tipAsDecimal * price
	totalBill := amountOfTip + price

	fmt.Printf("\nDollar amount of your tip: $%.2f\n", amountOfTip)
	fmt.Printf("Total bill after tip: $%.2f\n", totalBill)

}

// Layaway Computation: Given a total price and a number of months, compute the cost per month to buy the item on "layaway"
func layawayComputation(price float64, months int) {
	fmt.Println("Layaway Computation")
	fmt.Println("Price: ", price)
	fmt.Println("Months: ", months)

	costPerMonth := price / float64(months)

	fmt.Printf("Cost Per Month: $%.2f\n", costPerMonth)
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
