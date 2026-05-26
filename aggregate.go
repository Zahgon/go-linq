package linq

// Aggregate applies an accumulator function over a sequence.
//
// Aggregate method makes it simple to perform a calculation over a sequence of
// values. This method works by calling f() one time for each element in a source
// except the first one. Each time f() is called, Aggregate passes both the
// element from the sequence and an aggregated value (as the first argument to
// f()). The first element of the source is used as the initial aggregate value. The
// result of f() replaces the previous aggregated value.
//
// Aggregate returns the final result of f().
func (q Query) Aggregate(f func(accumulator, item any) any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// AggregateT is the typed version of Aggregate.
//
//   - f is of type: func(TSource, TSource) TSource
//
// NOTE: Aggregate has better performance than AggregateT.
func (q Query) AggregateT(f any) any { _ = "STUB: not implemented"; return *new(any) }

// AggregateWithSeed applies an accumulator function over a sequence. The
// specified seed value is used as the initial accumulator value.
//
// Aggregate method makes it simple to perform a calculation over a sequence of
// values. This method works by calling f() one time for each element in a source
// except the first one. Each time f() is called, Aggregate passes both the
// element from the sequence and an aggregated value (as the first argument to
// f()). The value of the seed parameter is used as the initial aggregate value.
// The result of f() replaces the previous aggregated value.
//
// Aggregate returns the final result of f().
func (q Query) AggregateWithSeed(seed any,
	f func(accumulator, item any) any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// AggregateWithSeedT is the typed version of AggregateWithSeed.
//
//   - f is of type "func(TAccumulate, TSource) TAccumulate"
//
// NOTE: AggregateWithSeed has better performance than
// AggregateWithSeedT.
func (q Query) AggregateWithSeedT(seed any,
	f any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// AggregateWithSeedBy applies an accumulator function over a sequence. The
// specified seed value is used as the initial accumulator value, and the
// specified function is used to select the result value.
//
// Aggregate method makes it simple to perform a calculation over a sequence of
// values. This method works by calling f() one time for each element in source.
// Each time func is called, Aggregate passes both the element from the sequence
// and an aggregated value (as the first argument to func). The value of the
// seed parameter is used as the initial aggregate value. The result of func
// replaces the previous aggregated value.
//
// The final result of func is passed to resultSelector to obtain the final
// result of Aggregate.
func (q Query) AggregateWithSeedBy(seed any,
	f func(accumulator, item any) any,
	resultSelector func(any) any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// AggregateWithSeedByT is the typed version of AggregateWithSeedBy.
//
//   - f is of type "func(TAccumulate, TSource) TAccumulate"
//   - resultSelectorFn is of type "func(TAccumulate) TResult"
//
// NOTE: AggregateWithSeedBy has better performance than
// AggregateWithSeedByT.
func (q Query) AggregateWithSeedByT(seed any,
	f any,
	resultSelectorFn any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
