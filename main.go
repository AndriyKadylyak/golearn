package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 2, 3, 4, 5}

	numsMax, numsMin := sumMaxMin(nums)

	fmt.Println(numsMax, numsMin)

}

func sumMaxMin(nums []int) (int, int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	min := nums[0] + nums[1] + nums[2] + nums[3]
	max := nums[1] + nums[2] + nums[3] + nums[4]

	return max, min
}
