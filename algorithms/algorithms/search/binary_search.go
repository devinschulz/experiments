package algorithms

func BinarySearch(arr []int, val int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == val {
			return mid
		} else if guess > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
