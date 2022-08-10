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
# Clone the repo
git clone https://github.com/Aerilate/sorting.git

# Go into the repo and run
cd sorting
go run main.go
~~~
~~~bash
# The output should look something like this:
$ go run main.go
Algorithm                Time (ms)
----------------------------------
MergesortConcurrent             66
QuicksortConcurrent             98
Quicksort                      116
LSDRadixsort                   135
StdLibSort                     139
Mergesort                      153
Heapsort                       320
~~~
