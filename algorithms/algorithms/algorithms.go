package algorithms

func BubbleSort(nums []int) []int {
	max := len(nums)
	for i := 0; i < max; i++ {
		for j := 0; j < max-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

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
