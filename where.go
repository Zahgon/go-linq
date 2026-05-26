package linq

// Where filters a collection of values based on a predicate.
func (q Query) Where(predicate func(any) bool) Query { _ = "STUB: not implemented"; return *new(Query) }

// WhereT is the typed version of Where.
//
//   - predicateFn is of type "func(TSource)bool"
//
// NOTE: Where has better performance than WhereT.
func (q Query) WhereT(predicateFn any) Query { _ = "STUB: not implemented"; return *new(Query) }

// WhereIndexed filters a collection of values based on a predicate. Each
// element's index is used in the logic of the predicate function.
//
// The first argument represents the zero-based index of the element within
// the collection. The second argument of predicate represents the element to test.
func (q Query) WhereIndexed(predicate func(int, any) bool) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// WhereIndexedT is the typed version of WhereIndexed.
//
//   - predicateFn is of type "func(int,TSource)bool"
//
// NOTE: WhereIndexed has better performance than WhereIndexedT.
func (q Query) WhereIndexedT(predicateFn any) Query { _ = "STUB: not implemented"; return *new(Query) }
