package sort

import (
	"sync"
)

func MergesortConcurrent(nums []int) {
	msc(nums, 0, len(nums)-1)
}

func msc(nums []int, lo int, hi int) {
	size := hi - lo + 1
	if size == 0 || size == 1 {
		return
	}
	if size < SWITCH_TO_INSERTION {
		Insertionsort(nums[lo : hi+1])
		return
	}

	median := lo + (hi-lo)/2

	if size > STOP_CONCURRENT_MERGESORT_AT_SIZE {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			msc(nums, lo, median)
		}()
		go func() {
			defer wg.Done()
			msc(nums, median+1, hi)
		}()
		wg.Wait()
	} else {
		msc(nums, lo, median)
		msc(nums, median+1, hi)
	}

	merge(nums, lo, median, median+1, hi)
}
