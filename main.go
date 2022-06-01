package main

import (
	"fmt"

	"github.com/Aerilate/sorting/benchmark"
	"github.com/Aerilate/sorting/sort"
)

const SIZE = 1000000

func main() {
	benchmark := benchmark.Benchmarker[[]int]{
		Funcs: []func([]int){
			sort.Heapsort,
			sort.Mergesort,
			sort.Quicksort,
			sort.MergesortConcurrent,
			sort.QuicksortConcurrent,
			sort.Radixsort,
		},
		InputGenerator:  prepareInput(),
		OutputValidator: isSorted,
	}
	fmt.Println(benchmark.Run().String())
}

func prepareInput() func() []int {
	var sorted []int
	for i := 0; i < SIZE; i++ {
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
