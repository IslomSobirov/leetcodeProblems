package problems

func searchInsert(nums []int, target int) int {
	j := 0
	for i := range nums {
		if nums[i] == target {
			return i
		}
		if nums[i] < target {
			j++
		}
	}
	return j
}
