package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sum(numbers...)
	fmt.Println("Sum of", numbers, "is", sum)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
