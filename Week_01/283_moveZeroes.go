package Week_01

//leetcode 283
//移动零
func moveZeroes(nums []int) {
	s := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[s] = nums[s], nums[i]
			s++
		}
	}
}
