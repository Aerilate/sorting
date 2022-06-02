package sort

const MIN_MERGE = 32

func Timsort(nums []int) {
	n := len(nums)
	minRun := calcMinRun(n)

	for start := 0; start < n; start += minRun {
		end := min(start+minRun-1, n-1)
		Insertionsort(nums[start : end+1])
	}

	size := minRun
	for size < n {
		for left := 0; left < n; left += 2 * size {
			mid := min(n-1, left+size-1)
			right := min(left+2*size-1, n-1)

			if mid < right {
				timMerge(nums, left, mid, right)
			}
		}
		size *= 2
	}
}

func timMerge(nums []int, lo int, med int, hi int) {
	len1, len2 := med-lo+1, hi-med
	var (
		left  []int
		right []int
	)
	for i := 0; i < len1; i++ {
		left = append(left, nums[lo+i])
	}
	for i := 0; i < len2; i++ {
		right = append(right, nums[med+1+i])
	}

	i, j, k := 0, 0, lo

	for i < len1 && j < len2 {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {
		nums[k] = left[i]
		k++
		i++
	}
	for j < len2 {
		nums[k] = right[j]
		k++
		j++
	}
}

func calcMinRun(n int) int {
	r := 0
	for n >= MIN_MERGE {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
