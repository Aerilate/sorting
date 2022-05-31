package sort

import "sync"

// At small inputs, the overhead for concurrency primitives is not worth it anymore
const STOP_CONCURRENT_QUICKSORT_AT_SIZE = 1000

func QuicksortConcurrent(nums []int) {
	qsc(nums, 0, len(nums)-1)
}

func qsc(nums []int, lo int, hi int) {
	if lo < 0 || hi < 0 || hi <= lo {
		return
	}

	pivot := partition(nums, lo, hi)
	if hi-lo+1 > STOP_CONCURRENT_QUICKSORT_AT_SIZE {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			qs(nums, lo, pivot)
		}()
		go func() {
			defer wg.Done()
			qs(nums, pivot+1, hi)
		}()
		wg.Wait()
	} else {
		qs(nums, lo, pivot)
		qs(nums, pivot+1, hi)
	}
}
