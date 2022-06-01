package benchmark

import (
	"fmt"
	"sort"
	"strings"
)

type ResultRow struct {
	Func    string
	Elapsed int64
}

type Result struct {
	rows []ResultRow
}

func (r Result) SortByFunc() Result {
	sort.Slice(r.rows, func(i, j int) bool {
		return r.rows[i].Func < r.rows[j].Func
	})
	return r
}

func (r Result) SortByElapsed() Result {
	sort.Slice(r.rows, func(i, j int) bool {
		return r.rows[i].Elapsed < r.rows[j].Elapsed
	})
	return r
}

func (r Result) String() string {
	var longestFuncName int
	for _, row := range r.rows {
		if len(row.Func) > longestFuncName {
			longestFuncName = len(row.Func)
		}
	}
	longestFuncName += 5

	strs := []string{fmt.Sprintf("%-*v %v", longestFuncName, "Algorithm", "Time (ms)")}
	strs = append(strs, strings.Repeat("-", len(strs[0])))

	for _, row := range r.rows {
		time := "-"
		if row.Elapsed >= 0 {
			time = fmt.Sprintf("%v", row.Elapsed)
		}
		strs = append(strs, fmt.Sprintf("%-*v %9v", longestFuncName, row.Func, time))
	}
	return strings.Join(strs, "\n")
}
