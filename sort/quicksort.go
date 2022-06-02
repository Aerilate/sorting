package sort

import "math/rand"

func Quicksort(nums []int) {
	qs(nums, 0, len(nums)-1)
}

func qs(nums []int, lo int, hi int) {
	if lo < 0 || hi < 0 || hi <= lo {
		return
	}
	if hi-lo < 50 {
		Insertionsort(nums[lo : hi+1])
		return
	}

	pivot := partition(nums, lo, hi)
	qs(nums, lo, pivot)
	qs(nums, pivot+1, hi)
}

// Hoare's partition
func partition(nums []int, lo int, hi int) int {
	pivot := rand.Intn(hi-lo+1) + lo
	i := lo - 1
	j := hi + 1

	for {
		i++
		for nums[i] < pivot {
			i++
		}

		j--
		for nums[j] > pivot {
			j--
		}

		if j <= i {
			return j
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}
