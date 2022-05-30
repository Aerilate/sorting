package main

import (
	"fmt"
	"math/rand"

	"github.com/Aerilate/sorting/sort"
)

const SIZE = 10

// Fisher-Yates shuffle
func Shuffle(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		j := rand.Intn(len(nums)-i) + i
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func isSorted(nums []int) bool {
	for idx, val := range nums {
		if idx != val {
			return false
		}
	}
	return true
}

func main() {
	var sorted []int
	for i := 0; i < SIZE; i++ {
		sorted = append(sorted, i)
	}

	unsorted := make([]int, len(sorted))
	copy(unsorted, sorted)
	Shuffle(unsorted)

	sort.Heapsort(unsorted)

	fmt.Println(unsorted)
	fmt.Println(sorted)
}
