package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/Aerilate/sorting/sort"
)

const SIZE = 1000000

func main() {
	var sorted []int
	for i := 0; i < SIZE; i++ {
		sorted = append(sorted, i)
	}

	type sortingAlgo func ([]int)
	algos := []sortingAlgo{
		sort.Heapsort, 
		sort.Mergesort, 
		sort.Quicksort,
		sort.MergesortConcurrent,
		sort.QuicksortConcurrent,
	}

	fmt.Printf("%-20v %-10v %v\n\n", "Algorithm", "Time (ms)", "Error")
	for _, algo := range algos {
		unsorted := make([]int, len(sorted))
		copy(unsorted, sorted)
		Shuffle(unsorted)

		elapsed, err := benchmark(algo, unsorted)
		fmt.Printf("%-20v %-10v %v\n", getFuncName(algo), elapsed, err)
	}
}

// Fisher-Yates shuffle
func Shuffle(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		j := rand.Intn(len(nums)-i) + i
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func validated(nums []int) bool {
	for idx, val := range nums {
		if idx != val {
			return false
		}
	}
	return true
}

func benchmark(sort func ([]int), input []int) (elapsed int64, err error) {
	start := time.Now()
	sort(input)
	elapsed = time.Since(start).Milliseconds()

	if !validated(input) {
		return elapsed, fmt.Errorf("algorithm did not sort correctly")
	}
	return elapsed, nil
}

func getFuncName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
    return strings.Split(name, ".")[2]
}
