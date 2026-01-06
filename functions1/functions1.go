package functions1

import "fmt"

type intTransformer func(int) int

func functions1() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Numbers: ", numbers)

	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled: ", doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled: ", tripled)

	sqred := transformNumbers(&numbers, sqr)
	fmt.Println("Squared: ", sqred)
}

func transformNumbers(numbers *[]int, transformation intTransformer) []int {
	doubled := []int{}
	for _, number := range *numbers {
		doubled = append(doubled, transformation(number))
	}
	return doubled
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func sqr(num int) int {
	return num * num
}
