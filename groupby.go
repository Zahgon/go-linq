package linq

// Group is a type used to store the result of GroupBy method.
type Group struct {
	Key   any
	Group []any
}

// GroupBy method groups the elements of a collection according to a specified
// key selector function and projects the elements for each group by using a
// specified function.
func (q Query) GroupBy(keySelector func(any) any,
	elementSelector func(any) any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// GroupByT is the typed version of GroupBy.
//
//   - keySelectorFn is of type "func(TSource) TKey"
//   - elementSelectorFn is of type "func(TSource) TElement"
//
// NOTE: GroupBy has better performance than GroupByT.
func (q Query) GroupByT(keySelectorFn any,
	elementSelectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
