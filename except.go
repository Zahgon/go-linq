package linq

// Except produces the set difference of two sequences. The set difference is
// the members of the first sequence that don't appear in the second sequence.
func (q Query) Except(q2 Query) Query { _ = "STUB: not implemented"; return *new(Query) }

// ExceptBy invokes a transform function on each element of a collection and
// produces the set difference of two sequences. The set difference is the
// members of the first sequence that don't appear in the second sequence.
func (q Query) ExceptBy(q2 Query, selector func(any) any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// ExceptByT is the typed version of ExceptBy.
//
//   - selectorFn is of type "func(TSource) TSource"
//
// NOTE: ExceptBy has better performance than ExceptByT.
func (q Query) ExceptByT(q2 Query,
	selectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
