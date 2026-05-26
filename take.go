package linq

// Take returns a specified number of contiguous elements from the start of a
// collection.
func (q Query) Take(count int) Query { _ = "STUB: not implemented"; return *new(Query) }

// TakeWhile returns elements from a collection as long as a specified condition
// is true and then skips the remaining elements.
func (q Query) TakeWhile(predicate func(any) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// TakeWhileT is the typed version of TakeWhile.
//
//   - predicateFn is of type "func(TSource)bool"
//
// NOTE: TakeWhile has better performance than TakeWhileT.
func (q Query) TakeWhileT(predicateFn any) Query { _ = "STUB: not implemented"; return *new(Query) }

// TakeWhileIndexed returns elements from a collection as long as a specified
// condition is true. The element's index is used in the logic of the predicate
// function. The first argument of predicate represents the zero-based index of
// the element within the collection. The second argument represents the element to
// test.
func (q Query) TakeWhileIndexed(predicate func(int, any) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// TakeWhileIndexedT is the typed version of TakeWhileIndexed.
//
//   - predicateFn is of type "func(int,TSource)bool"
//
// NOTE: TakeWhileIndexed has better performance than TakeWhileIndexedT.
func (q Query) TakeWhileIndexedT(predicateFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
