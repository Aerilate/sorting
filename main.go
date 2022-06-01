package main

import (
	"fmt"

	"github.com/Aerilate/sorting/benchmark"
	"github.com/Aerilate/sorting/sort"
)

const INPUT_SIZE = 1000000

func main() {
	benchmark := benchmark.Benchmarker[[]int]{
		Funcs: []func([]int){
			// sort.NoSort,
			sort.StdLibSort,
			sort.Heapsort,
			sort.Mergesort,
			sort.Quicksort,
			sort.MergesortConcurrent,
			sort.QuicksortConcurrent,
			sort.Radixsort,
		},
		InputGenerator:  shuffleInputFunc(),
		OutputValidator: isSorted,
	}
	fmt.Println(benchmark.Run().SortByElapsed().String())
}

func shuffleInputFunc() func() []int {
	var sorted []int
	for i := 0; i < INPUT_SIZE; i++ {
		sorted = append(sorted, i)
	}

	return func() []int {
		unsorted := make([]int, len(sorted))
		copy(unsorted, sorted)
		sort.Shuffle(unsorted)
		return unsorted
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
