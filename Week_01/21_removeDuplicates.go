package Week_01

//leetcode 21
//删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	i := 0
	cnt := len(nums)
	if cnt < 2 {
		return 1
	}
	for j := 1; j < cnt; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
