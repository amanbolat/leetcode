package solution

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	halfLen := totalLen / 2
	mod := totalLen % 2
	poss := make(map[int]bool)

	if mod == 0 {
		poss[halfLen] = true
		poss[halfLen-1] = true
	} else {
		poss[halfLen] = true
	}

	var i1, i2 int

	var sum float64
	for i := 0; i <= halfLen; i++ {
		var cur int
		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] < nums2[i2] {
				cur = nums1[i1]
				i1++
			} else {
				cur = nums2[i2]
				i2++
			}
		} else if i1 < len(nums1) {
			cur = nums1[i1]
			i1++
		} else {
			cur = nums2[i2]
			i2++
		}
		if _, ok := poss[i]; ok {
			sum += float64(cur)
		}
	}

	return sum/float64(len(poss))
}