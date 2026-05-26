package linq

// Join correlates the elements of two collections based on matching keys.
//
// A join refers to the operation of correlating the elements of two sources of
// information based on a common key. Join brings the two information sources
// and the keys by which they are matched together in one method call. This
// differs from the use of SelectMany, which requires more than one method call
// to perform the same operation.
//
// Join preserves the order of the elements of outer collection, and for each of
// these elements, the order of the matching elements of inner.
func (q Query) Join(inner Query,
	outerKeySelector func(any) any,
	innerKeySelector func(any) any,
	resultSelector func(outer any, inner any) any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// JoinT is the typed version of Join.
//
//   - outerKeySelectorFn is of type "func(TOuter) TKey"
//   - innerKeySelectorFn is of type "func(TInner) TKey"
//   - resultSelectorFn is of type "func(TOuter,TInner) TResult"
//
// NOTE: Join has better performance than JoinT.
func (q Query) JoinT(inner Query,
	outerKeySelectorFn any,
	innerKeySelectorFn any,
	resultSelectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
