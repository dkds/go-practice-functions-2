package main

import "fmt"

func main() {
	number := 6
	factorial := factorial(number)
	fmt.Println("Factorial of", number, "is", factorial)
}

// func factorial(number int) int {
// 	total := 1
// 	for i := 1; i <= number; i++ {
// 		total *= i
// 	}
// 	return total
// }

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
