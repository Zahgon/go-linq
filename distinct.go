package linq

// Distinct method returns distinct elements from a collection. The result is an
// unordered collection that contains no duplicate values.
func (q Query) Distinct() Query { _ = "STUB: not implemented"; return *new(Query) }

// Distinct method returns distinct elements from a collection. The result is an
// ordered collection that contains no duplicate values.
//
// NOTE: Distinct method on OrderedQuery type has better performance than
// Distinct method on Query type.
func (oq OrderedQuery) Distinct() OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// DistinctBy method returns distinct elements from a collection. This method
// executes selector function for each element to determine a value to compare.
// The result is an unordered collection that contains no duplicate values.
func (q Query) DistinctBy(selector func(any) any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// DistinctByT is the typed version of DistinctBy.
//
//   - selectorFn is of type "func(TSource) TSource".
//
// NOTE: DistinctBy has better performance than DistinctByT.
func (q Query) DistinctByT(selectorFn any) Query { _ = "STUB: not implemented"; return *new(Query) }
