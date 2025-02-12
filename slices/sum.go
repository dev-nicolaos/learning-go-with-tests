package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numberCollections ...[]int) []int {
	sums := make([]int, len(numberCollections))
	for i, collection := range numberCollections {
		sums[i] = Sum(collection)
	}
	return sums
}
