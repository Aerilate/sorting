package benchmark

import (
	"fmt"
	"strings"
)

type Benchmarker[T any] struct {
	Funcs               []func(T)
	InputGenerator  func() T
	OutputValidator func(T) bool
}

type Result struct {
	rows []ResultRow
}

type ResultRow struct {
	funcName string
	elapsed  int64
	valid    bool
}

func (b *Benchmarker[T]) Run() (result Result) {
	for _, f := range b.Funcs {
		input := b.InputGenerator()
		elapsed := Stopwatch(f, input)
		result.rows = append(result.rows, ResultRow{
			funcName: FuncName(f),
			elapsed:  elapsed,
			valid:    b.OutputValidator(input),
		})
	}
	return result
}

func (r Result) String() string {
	strs := []string{fmt.Sprintf("%-20v %-10v %v", "Algorithm", "Time (ms)", "Correct")}

	for _, row := range r.rows {
		s := fmt.Sprintf("%-20v %-10v %v", row.funcName, row.elapsed, row.valid)
		strs = append(strs, s)
	}
	return strings.Join(strs, "\n")
}
