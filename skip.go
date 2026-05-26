package linq

// Skip bypasses a specified number of elements in a collection and then returns
// the remaining elements.
func (q Query) Skip(count int) Query { _ = "STUB: not implemented"; return *new(Query) }

// SkipWhile bypasses elements in a collection as long as a specified condition
// is true and then returns the remaining elements.
//
// This method tests each element by using predicate and skips the element if
// the result is true. After the predicate function returns false for an
// element, that element and the remaining elements in source are returned and
// there are no more invocations of predicate.
func (q Query) SkipWhile(predicate func(any) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SkipWhileT is the typed version of SkipWhile.
//
//   - predicateFn is of type "func(TSource)bool"
//
// NOTE: SkipWhile has better performance than SkipWhileT.
func (q Query) SkipWhileT(predicateFn any) Query { _ = "STUB: not implemented"; return *new(Query) }

// SkipWhileIndexed bypasses elements in a collection as long as a specified
// condition is true and then returns the remaining elements. The element's
// index is used in the logic of the predicate function.
//
// This method tests each element by using predicate and skips the element if
// the result is true. After the predicate function returns false for an
// element, that element and the remaining elements in source are returned and
// there are no more invocations of predicate.
func (q Query) SkipWhileIndexed(predicate func(int, any) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SkipWhileIndexedT is the typed version of SkipWhileIndexed.
//
//   - predicateFn is of type "func(int,TSource)bool"
//
// NOTE: SkipWhileIndexed has better performance than SkipWhileIndexedT.
func (q Query) SkipWhileIndexedT(predicateFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
