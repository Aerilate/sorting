package benchmark

import (
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Stopwatch[T any](f func(T), input T) (elapsed int64) {
	start := time.Now()
	f(input)
	return time.Since(start).Milliseconds()
}

// FuncName returns a string representing a function's name in the code
func FuncName[T any](fn func(T)) string {
	pc := reflect.ValueOf(fn).Pointer()
	name := runtime.FuncForPC(pc).Name()
	splitName := strings.Split(name, ".")
	return splitName[len(splitName)-1]
}
