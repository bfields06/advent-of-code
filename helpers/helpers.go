package helpers

func SumCollection(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}