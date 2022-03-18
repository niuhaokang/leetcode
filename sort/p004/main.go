package main

import "sort"

func singleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i:=0; i<len(nums); i+=3 {
		if i == len(nums)-1 {return nums[i]}
		if nums[i] != nums[i+1] {return nums[i]}
	}
	return -1
}