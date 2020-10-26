package problems

func search(nums []int, target int) int {
	for i, j := range nums {
		if j == target {
			return i
		}
	}

	return -1
}
