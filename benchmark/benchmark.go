package benchmark

type Benchmarker[T any] struct {
	Funcs           []func(T)
	InputGenerator  func() T
	OutputValidator func(T) bool
}

func (b *Benchmarker[T]) Run() (res Result) {
	for _, f := range b.Funcs {
		input := b.InputGenerator()
		elapsed := Stopwatch(f, input)
		if !b.OutputValidator(input) {
			elapsed = -1
		}

		res.rows = append(res.rows, ResultRow{
			Func:    FuncName(f),
			Elapsed: elapsed,
		})
	}
	return res
}
