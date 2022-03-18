package main

func numSubarrayProductLessThanK(nums []int, k int) (cnt int) {
	window := 1
	for left, right:=0, 0; right<len(nums); right++ {
		window *= nums[right]
		for left<=right && window >= k {
			window /= nums[left]
			left++
		}
		cnt += right - left + 1
	}
	return
}