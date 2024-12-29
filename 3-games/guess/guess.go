package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	fmt.Println("Please think of a number btwn 1 and 100")
	fmt.Println("Press enter when ready")
	scanner.Scan()

	for {
		if (low > high) {
			fmt.Println("You are cheating!")
			break
		}
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that: ")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()

		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Invalid response, try again")
		}
	}
}