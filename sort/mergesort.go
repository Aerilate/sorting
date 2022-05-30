package sort

func Mergesort(nums []int) {
	ms(nums, 0, len(nums)-1)
}

func ms(nums []int, lo int, hi int) {
	size := hi - lo + 1
	if size == 0 || size == 1 {
		return
	}

	median := lo + (hi-lo)/2
	ms(nums, lo, median)
	ms(nums, median+1, hi)
	merge(nums, lo, median, median+1, hi)
}

func merge(arr []int, lo1 int, hi1 int, lo2 int, hi2 int) {
	size := hi2 - lo1 + 1
	arrCopy := make([]int, size)
	
	curr1 := lo1
	curr2 := lo2

	for i := 0; i < size; i++ {
		if curr1 > hi1 && curr2 > hi2 { // both subarrays done
			break
		}

		if curr1 <= hi1 && (curr2 <= hi2 && arr[curr1] <= arr[curr2] || curr2 > hi2) {
			arrCopy[i] = arr[curr1]
			curr1++
		} else {
			arrCopy[i] = arr[curr2]
			curr2++
		}
	}

	// copy arrCopy into arr
	for i := 0; i < size; i++ {
		arr[i+lo1] = arrCopy[i]
	}
}
