package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	// sum of nums[0,i)
	preSum := make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] += preSum[i] + num
	}

	shortestLen := math.MaxInt32
	for i:=1; i<=len(nums); i++ {
		for j:=0; j<i; j++ {
			if preSum[i] - preSum[j] >= target {
				shortestLen = min(shortestLen, i - j)
			}
		}
	}

	if shortestLen == math.MaxInt32 {
		return 0
	}
	return shortestLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}