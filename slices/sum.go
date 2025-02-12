package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numberCollections ...[]int) (sums []int) {
	for _, collection := range numberCollections {
		sums = append(sums, Sum(collection))
	}
	return
}

func SumAllTails(numberCollections ...[]int) (sums []int) {
	for _, collection := range numberCollections {
		sums = append(sums, Sum(collection[1:]))
	}
	return
}
