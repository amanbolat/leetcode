package solution_test

import (
	"fmt"
	solution "github.com/amanbolat/leetcode/hard-median-of-two-sorted-arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct{
		Nums1    []int
		Nums2    []int
		Expected float64
	}{
		{Nums1: []int{1,3}, Nums2: []int{2}, Expected: 2},
		{Nums1: []int{1,2}, Nums2: []int{3,4}, Expected: 2.5},
		{Nums1: []int{0,0}, Nums2: []int{0,0}, Expected: 0},
		{Nums1: []int{}, Nums2: []int{1}, Expected: 1},
		{Nums1: []int{2}, Nums2: []int{}, Expected: 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("nums1: %v, nums2: %v, expected: %v",
			tc.Nums1,
			tc.Nums2,
			tc.Expected),
			func(t *testing.T) {
			assert.Equal(t, tc.Expected, solution.FindMedianSortedArrays(tc.Nums1, tc.Nums2))
		})
	}
}
