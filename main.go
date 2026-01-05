package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	doubled := []int{}
	for number := range *numbers {
		doubled = append(doubled, number*2)
	}
	return doubled
}
