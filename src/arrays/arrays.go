package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers)) //Allocates a zeroed array
	for i, nums := range numbers {
		fmt.Println(i, numbers)
		sums[i] = Sum(nums)
	}
	return sums

}
