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
