package sort

// At small inputs, the overhead for concurrency primitives is not worth it anymore
const STOP_CONCURRENT_MERGESORT_AT_SIZE = 1000
const STOP_CONCURRENT_QUICKSORT_AT_SIZE = 1000

const SWITCH_TO_INSERTION = 50
