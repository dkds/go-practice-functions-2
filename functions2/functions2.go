package functions2

import "fmt"

type intTransformer func(int) int

func functions2() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Numbers: ", numbers)

	doubleTransformer := createTransformer(2)
	tripleTransformer := createTransformer(3)

	doubled := transformNumbers(&numbers, doubleTransformer)
	fmt.Println("Doubled: ", doubled)

	tripled := transformNumbers(&numbers, tripleTransformer)
	fmt.Println("Tripled: ", tripled)

	sqred := transformNumbers(&numbers, func(i int) int { return i * i })
	fmt.Println("Squared: ", sqred)
}

func transformNumbers(numbers *[]int, transformation intTransformer) []int {
	doubled := []int{}
	for _, number := range *numbers {
		doubled = append(doubled, transformation(number))
	}
	return doubled
}

func createTransformer(factor int) intTransformer {
	return func(i int) int {
		return i * factor
	}
}
