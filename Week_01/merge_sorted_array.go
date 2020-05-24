// 88. 合并两个有序数组
// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1 := m-1
	n2 := n-1
	l := m+n-1
	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums1[n2] {
			nums1[l] = nums1[n1]
			n1--
		} else {
			nums1[l] = nums2[n2]
			n2--
		}
		l--
	}
}