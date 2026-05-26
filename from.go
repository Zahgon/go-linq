package linq

import (
	"context"
	"iter"
)

// Query is the type returned from query functions. It can be iterated manually
// as shown in the example.
type Query struct {
	Iterate iter.Seq[any]
}

// KeyValue is a type used to iterate over a map. This type is also used by ToMap()
// method to output the result of a query into a map.
type KeyValue struct {
	Key   any
	Value any
}

// Iterable is an interface that has to be implemented by a custom collection
// to work with linq.
type Iterable interface {
	Iterate() iter.Seq[any]
}

// FromSlice initializes a linq query with a passed slice.
func FromSlice[S ~[]T, T any](source S) Query { _ = "STUB: not implemented"; return *new(Query) }

// FromMap initializes a linq query with a passed map.
func FromMap[M ~map[K]V, K comparable, V any](source M) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// FromChannel initializes a linq query with a passed channel, linq iterates over
// the channel until it is closed.
func FromChannel[T any](source <-chan T) Query { _ = "STUB: not implemented"; return *new(Query) }

// FromChannelWithContext initializes a linq query with a passed channel
// and stops iterating either when the channel is closed or when the context is canceled.
func FromChannelWithContext[T any](ctx context.Context, source <-chan T) Query {
	_ = "STUB: not implemented"
	return *new(Query)
}

// Context canceled or deadline exceeded

// Channel closed or Consumer stopped early

// FromString initializes a query from a string, iterating over its runes.
func FromString[S ~string](source S) Query { _ = "STUB: not implemented"; return *new(Query) }

// FromIterable initializes a linq query with a custom collection passed. This
// collection has to implement Iterable.
func FromIterable(source Iterable) Query { _ = "STUB: not implemented"; return *new(Query) }

// From initializes a Query from a supported data source by inspecting its
// type at runtime. It panics if the source type is not supported.
//
// NOTE: It is recommended to call the specific From* function directly
// (e.g., FromSlice, FromMap, etc.). This unified function is less efficient
// because it relies on runtime reflection.
func From(source any) Query { _ = "STUB: not implemented"; return *new(Query) }

// Range generates a sequence of integral numbers within a specified range.
func Range(start, count int) Query { _ = "STUB: not implemented"; return *new(Query) }

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](value T, count int) Query { _ = "STUB: not implemented"; return *new(Query) }
