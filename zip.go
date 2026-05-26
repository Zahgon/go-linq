package linq

// Zip applies a specified function to the corresponding elements of two
// collections, producing a collection of the results.
//
// The method steps through the two input collections, applying function
// resultSelector to corresponding elements of the two collections. The method
// returns a collection of the values that are returned by resultSelector. If
// the input collections do not have the same number of elements, the method
// combines elements until it reaches the end of one of the collections. For
// example, if one collection has three elements and the other one has four, the
// result collection has only three elements.
func (q Query) Zip(q2 Query,
	resultSelector func(any, any) any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// ZipT is the typed version of Zip.
//
//   - resultSelectorFn is of type "func(TFirst,TSecond)TResult"
//
// NOTE: Zip has better performance than ZipT.
func (q Query) ZipT(q2 Query,
	resultSelectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
