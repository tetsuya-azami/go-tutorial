package mylib

func Average(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total / len(nums)
}
