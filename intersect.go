package linq

// Intersect produces the set intersection of the source collection and the
// provided input collection. The intersection of two sets A and B is defined as
// the set that contains all the elements of A that also appear in B, but no
// other elements.
func (q Query) Intersect(q2 Query) Query { _ = "STUB: not implemented"; return *new(Query) }

// IntersectBy produces the set intersection of the source collection and the
// provided input collection. The intersection of two sets A and B is defined as
// the set that contains all the elements of A that also appear in B, but no
// other elements.
//
// IntersectBy invokes a transform function on each element of both collections.
func (q Query) IntersectBy(q2 Query,
	selector func(any) any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// IntersectByT is the typed version of IntersectBy.
//
//   - selectorFn is of type "func(TSource) TSource"
//
// NOTE: IntersectBy has better performance than IntersectByT.
func (q Query) IntersectByT(q2 Query,
	selectorFn any) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}
