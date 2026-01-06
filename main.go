package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum is", sum)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
