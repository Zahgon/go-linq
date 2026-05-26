package linq

// Append inserts an item to the end of a collection, so it becomes the last
// item.
func (q Query) Append(item any) Query { _ = "STUB: not implemented"; return *new(Query) }

// Concat concatenates two collections.
//
// The Concat method differs from the Union method because the Concat method
// returns all the original elements in the input sequences. The Union method
// returns only unique elements.
func (q Query) Concat(q2 Query) Query { _ = "STUB: not implemented"; return *new(Query) }

// Prepend inserts an item to the beginning of a collection, so it becomes the
// first item.
func (q Query) Prepend(item any) Query { _ = "STUB: not implemented"; return *new(Query) }
