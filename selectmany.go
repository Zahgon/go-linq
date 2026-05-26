package linq

// SelectMany projects each element of a collection to a Query, iterates and
// flattens the resulting collection into one collection.
func (q Query) SelectMany(selector func(any) Query) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SelectManyT is the typed version of SelectMany.
//
//   - selectorFn is of type "func(TSource)Query"
//
// NOTE: SelectMany has better performance than SelectManyT.
func (q Query) SelectManyT(selectorFn any) Query { _ = "STUB: not implemented"; return *new(Query) }

// SelectManyIndexed projects each element of a collection to a Query, iterates
// and flattens the resulting collection into one collection.
//
// The first argument to selector represents the zero-based index of that
// element in the source collection. This can be useful if the elements are in a
// known order and you want to do something with an element at a particular
// index, for example. It can also be useful if you want to retrieve the index
// of one or more elements. The second argument to selector represents the
// element to process.
func (q Query) SelectManyIndexed(selector func(index int, outer any) Query) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SelectManyIndexedT is the typed version of SelectManyIndexed.
//
//   - selectorFn is of type "func(int,TSource)Query"
//
// NOTE: SelectManyIndexed has better performance than SelectManyIndexedT.
func (q Query) SelectManyIndexedT(selectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SelectManyBy projects each element of a collection to a Query, iterates and
// flattens the resulting collection into one collection, and invokes a result
// selector function on each element therein.
func (q Query) SelectManyBy(
	selector func(outer any) Query,
	resultSelector func(inner, outer any) any,
) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SelectManyByT is the typed version of SelectManyBy.
//
//   - selectorFn is of type "func(TSource)Query"
//   - resultSelectorFn is of type "func(TSource,TCollection)TResult"
//
// NOTE: SelectManyBy has better performance than SelectManyByT.
func (q Query) SelectManyByT(selectorFn any,
	resultSelectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SelectManyByIndexed projects each element of a collection to a Query,
// iterates and flattens the resulting collection into one collection, and
// invokes a result selector function on each element therein. The index of each
// source element is used in the intermediate projected form of that element.
func (q Query) SelectManyByIndexed(
	selector func(index int, outer any) Query,
	resultSelector func(inner, outer any) any,
) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// SelectManyByIndexedT is the typed version of SelectManyByIndexed.
//
//   - selectorFn is of type "func(int,TSource)Query"
//   - resultSelectorFn is of type "func(TSource,TCollection)TResult"
//
// NOTE: SelectManyByIndexed has better performance than
// SelectManyByIndexedT.
func (q Query) SelectManyByIndexedT(selectorFn any,
	resultSelectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
