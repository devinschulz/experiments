package algorithms

func InsertionSort(nums []int) []int {
	for i, num := range nums {
		j := i - 1
		for ; j >= 0 && nums[j] > num; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = num
	}
	return nums
}
