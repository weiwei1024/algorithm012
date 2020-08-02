package src

//搜索旋转排序数组
//有序 可以使用二分
//关键在于判断哪部分是有序的
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l != r {
		mid = l + (r-l)/2 //防止越界
		if mid == target {
			return mid
		}
		//左侧有序
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//右侧有序
			if nums[mid] <= target && target < nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}
