package binPointer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0)
	for i:=0; i<len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for ;left<right&&nums[left]==nums[left-1]; left++ {}
				for ;left<right&&nums[right]==nums[right+1]; right-- {}
			} else if nums[i] + nums[left] + nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
