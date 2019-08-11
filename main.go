package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get user input as length of prime list
	userInput := bufio.NewScanner(os.Stdin)

	// when user enters the number
	for userInput.Scan()  {
		// convert input string to integer
		listNumber, _ := strconv.Atoi(userInput.Text())

		// iterate from 1 to n
		for i := 1; i < listNumber; i++ {
			// only prints prime
			if findPrime(i) {
				fmt.Print(i)
			}
		}
	}
}

func findPrime(n int) bool {
	// prime cannot be below 1
	if n <= 1 {
		return false
	}

	// prime cannot be divided with other number than itself
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}