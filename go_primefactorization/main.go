package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func readPrimeFromUser() int {
	args := os.Args

	// Check that there is only 1 user command line argument
	if len(args) != 2 {
		log.Fatal("Usage:", args[0], " Integer")
	}

	// If there is one, check that it is an integer.
	integer, err := strconv.Atoi(args[1])

	// If it isn't, also exit.
	if err != nil {
		log.Fatal(err)
	}

	if integer < 0 {
		log.Fatal("Input number must be a positive integer")
	}

	// If it is, return that as an int. Or string? I'll figure it out.
	return integer
}

func primeSieve(number int) []int {
	numbers := make([]int, number-1)
	isPrime := make([]bool, number-1)

	for i := 2; i <= number; i++ {
		numbers[i-2] = i
		isPrime[i-2] = true
	}

	for i := 0; i < len(numbers); i++ {

		currNum := numbers[i]

		for j := i + 1; j < len(numbers); j++ {

			if (numbers[j] % currNum) == 0 {
				isPrime[j] = false
			}
		}
	}

	primeNums := make([]int, 0)

	for i := 0; i < len(numbers); i++ {

		if isPrime[i] {
			primeNums = append(primeNums, numbers[i])
		}
	}

	return (primeNums)
}

func primeFactorizationOfFactorial(number int) {

	// Get prime numbers up until this number
	primeNums := primeSieve(number)

	numAppearance := 0

	// Iterate over that list and find out how many times it appears as a factor
	for i := 0; i < len(primeNums); i++ {
		numAppearance = countPrimeFactorAppearance(number, primeNums[i])

		fmt.Println(primeNums[i], "^", numAppearance)
	}

	// Convert that count to a string in the form "n^(number of appearances), "

}

func countPrimeFactorAppearance(number int, factor int) int {

	numAppearance := 0

	divisor := factor
	quotient := number / divisor

	for quotient != 0 {
		numAppearance += quotient

		divisor *= factor

		quotient = number / divisor
	}

	return numAppearance
}

func main() {

	integer := readPrimeFromUser()

	primeFactorizationOfFactorial(integer)

}
