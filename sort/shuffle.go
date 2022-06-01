package sort

import "math/rand"

// Fisher-Yates shuffle
func Shuffle(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		j := rand.Intn(len(nums)-i) + i
		nums[i], nums[j] = nums[j], nums[i]
	}
}
