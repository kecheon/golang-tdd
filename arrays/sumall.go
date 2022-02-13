package arrays

func SumAll(toSum ...[]int) (sums []int) {
	for _, numbers := range toSum {
		sums = append(sums, Sum(numbers))
	}
	return
}
