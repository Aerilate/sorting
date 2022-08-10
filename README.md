# :bar_chart: Sorting
A collection of sorting algorithms with a small benchmarking library, implemented in Go.
\
\
Some of the algorithms included are:
* quicksort (with and without concurrency)
* mergesort (with and without concurrency)
* heapsort
* LSD radix sort

## Goal
This project is an exercise using the Go standard library. Most of the creativity went into the benchmarking package and figuring out how to introduce Go-style concurrency to classic sorting algorithms.

## Usage
To run the benchmark yourself, you will need Go 1.18 installed.
~~~bash
# Run
go run main.go

# The output should look something like this:
Algorithm                Time (ms)
----------------------------------
MergesortConcurrent             27
LSDRadixsort                    30
QuicksortConcurrent             32
Quicksort                       52
Mergesort                       76
Timsort                        107
StdLibSort                     124
Heapsort                       275
~~~
