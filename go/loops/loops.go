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
func PrintNumbers() {
	fmt.Println("\nLooping through a list of numbers!")

	numberList := []int{17, 23, 53, 2, 50}
	for _, number := range numberList {
		fmt.Println(number)
	}

}

// Compute the sum and the average of a list of 5 numbers
func SumAndAvg() {
	fmt.Println("\nSum and Average of list")

	numberList := []int{17, 23, 53, 2, 50}

	var sum int
	for _, number := range numberList {
		sum += number
	}

	average := float64(sum) / float64(len(numberList))

	fmt.Println("SUM: ", sum)
	fmt.Printf("AVERAGE: %.2f\n", average)
}

// Given a list of 0 - 100
// If a number is divisible by 15 print fizzbuzz
// If a number is divisible by 5, print fizz
// If a number is divisible by 3, print buzz
func FizzBuzz() {
	fmt.Println("\nFIZZ BUZZ")
	fmt.Println(strings.Repeat("_", 30))

	numList := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
		61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
		91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}

	for _, number := range numList {
		if number%15 == 0 {
			fmt.Println("FizzBuzz", number)
		} else if number%5 == 0 {
			fmt.Println("Fizz", number)
		} else if number%3 == 0 {
			fmt.Println("Buzz", number)
		}
	}
}

// Given a list of 5 item prices
// For each price, compute the price with tax
// Print the total price along with each price with tax
func TaxCalculation() {
	fmt.Println("\nTAX CALCULATION")
	fmt.Println(strings.Repeat("_", 30))
	tax := GetValidatedNumber("Enter a number(0 - 100) for tax: ", 0, 100)
	taxDecimal := float64(tax) / 100

	numList := []int{42, 23, 5, 24, 29}
	var totalTax float64
	var total float64

	for i, number := range numList {
		numberTax := float64(number) * taxDecimal
		totalTax += numberTax
		total += float64(number)
		fmt.Printf("Tax for item %d: %.2f\n", i, numberTax)
	}

	fmt.Printf("Total Tax: $%.2f\n", totalTax)
	fmt.Printf("Full Price With Tax: $%.2f\n", total)

}

// Given a list of 5 numbers, print the list in ascending order
func Sorting() {

	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("\tBUBBLE SORT")
	fmt.Println(strings.Repeat("-", 55))

	nums := []int{5, 58, 42, 2, 38}
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("ORIGINAL LIST")
	fmt.Println(nums)
	fmt.Println(strings.Repeat("-", 55))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) - 1; j++ {
	
			if nums[j] > nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}

	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("SORTED LIST")
	fmt.Println(nums)
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
