package sort

import "fmt"

func LSDRadixsort(nums []int) {
	maxVal, err := max(nums)
	if err != nil {
		return
	}

	exp := 1
	for maxVal/exp > 1 {
		countingSort(nums, exp)
		exp *= 10
	}
}

func countingSort(nums []int, exp int) {
	output := make([]int, len(nums))
	var count [10]int

	for i := 0; i < len(nums); i++ {
		index := nums[i] / exp
		count[index%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		index := nums[i] / exp
		output[count[index%10]-1] = nums[i]
		count[index%10]--
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = output[i]
	}
}

func max(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("empty slice")
	}

	res := nums[0]
	for _, val := range nums {
		if val > res {
			res = val
		}
	}
	return res, nil
}
