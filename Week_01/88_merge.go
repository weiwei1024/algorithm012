package Week_01

//leetcode 88
//合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	c := m + n - 1
	for c >= 0 {
		if i < 0 {
			nums1[c] = nums2[j]
			j--
			c--
			continue
		}

		if j < 0 {
			nums1[c] = nums1[i]
			i--
			c--
			continue
		}

		if nums1[i] > nums2[j] {
			nums1[c] = nums1[i]
			i--
		} else {
			nums1[c] = nums2[j]
			j--
		}
		c--
	}
}
