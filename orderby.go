package linq

type order struct {
	selector func(any) any
	compare  comparer
	desc     bool
}

// OrderedQuery is the type returned from OrderBy, OrderByDescending ThenBy and
// ThenByDescending functions.
type OrderedQuery struct {
	Query
	original Query
	orders   []order
}

// OrderBy sorts the elements of a collection in ascending order. Elements are
// sorted according to a key.
func (q Query) OrderBy(selector func(any) any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// OrderByT is the typed version of OrderBy.
//
//   - selectorFn is of type "func(TSource) TKey"
//
// NOTE: OrderBy has better performance than OrderByT.
func (q Query) OrderByT(selectorFn any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// OrderByDescending sorts the elements of a collection in descending order.
// Elements are sorted according to a key.
func (q Query) OrderByDescending(selector func(any) any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// OrderByDescendingT is the typed version of OrderByDescending.
//   - selectorFn is of type "func(TSource) TKey"
//
// NOTE: OrderByDescending has better performance than OrderByDescendingT.
func (q Query) OrderByDescendingT(selectorFn any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// ThenBy performs a subsequent ordering of the elements in a collection in
// ascending order. This method enables you to specify multiple sort criteria by
// applying any number of ThenBy or ThenByDescending methods.
func (oq OrderedQuery) ThenBy(
	selector func(any) any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// ThenByT is the typed version of ThenBy.
//   - selectorFn is of type "func(TSource) TKey"
//
// NOTE: ThenBy has better performance than ThenByT.
func (oq OrderedQuery) ThenByT(selectorFn any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// ThenByDescending performs a subsequent ordering of the elements in a
// collection in descending order. This method enables you to specify multiple
// sort criteria by applying any number of ThenBy or ThenByDescending methods.
func (oq OrderedQuery) ThenByDescending(selector func(any) any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// ThenByDescendingT is the typed version of ThenByDescending.
//   - selectorFn is of type "func(TSource) TKey"
//
// NOTE: ThenByDescending has better performance than ThenByDescendingT.
func (oq OrderedQuery) ThenByDescendingT(selectorFn any) OrderedQuery {
	_ = "STUB: not implemented"
	return *new(OrderedQuery)
}

// Sort returns a new query by sorting elements with provided less function in
// ascending order. The comparer function should return true if the parameter i
// is less than j. While this method is uglier than chaining OrderBy,
// OrderByDescending, ThenBy and ThenByDescending methods, its performance is
// much better.
func (q Query) Sort(less func(i, j any) bool) Query { _ = "STUB: not implemented"; return *new(Query) }

// SortT is the typed version of Sort.
//   - lessFn is of type "func(TSource,TSource) bool"
//
// NOTE: Sort has better performance than SortT.
func (q Query) SortT(lessFn any) Query { _ = "STUB: not implemented"; return *new(Query) }

type sorter struct {
	items []any
	less  func(i, j any) bool
}

func (s sorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (s sorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s sorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (q Query) sort(orders []order) (r []any) { _ = "STUB: not implemented"; return nil }

func (q Query) lessSort(less func(i, j any) bool) (r []any) { _ = "STUB: not implemented"; return nil }
