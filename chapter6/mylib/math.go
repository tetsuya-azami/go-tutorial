/*
Package mylib provides custom math functions
*/
package mylib

// Average returns the average of a series of numbers
func Average(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total / len(nums)
}
