package linq

// All determines whether all elements of a collection satisfy a condition.
func (q Query) All(predicate func(any) bool) bool { _ = "STUB: not implemented"; return false }

// AllT is the typed version of All.
//
//   - predicateFn is of type "func(TSource) bool"
//
// NOTE: All has better performance than AllT.
func (q Query) AllT(predicateFn any) bool { _ = "STUB: not implemented"; return false }

// Any determines whether any element of a collection exists.
func (q Query) Any() bool { _ = "STUB: not implemented"; return false }

// AnyWith determines whether any element of a collection satisfies a condition.
func (q Query) AnyWith(predicate func(any) bool) bool { _ = "STUB: not implemented"; return false }

// AnyWithT is the typed version of AnyWith.
//
//   - predicateFn is of type "func(TSource) bool"
//
// NOTE: AnyWith has better performance than AnyWithT.
func (q Query) AnyWithT(predicateFn any) bool { _ = "STUB: not implemented"; return false }

// Average computes the average of a collection of numeric values.
// It panics if the sequence contains non-numeric types.
// It returns math.NaN() if the sequence is empty.
func (q Query) Average() (r float64) { _ = "STUB: not implemented"; return 0 }

// Contains determines whether a collection contains a specified element.
func (q Query) Contains(value any) bool { _ = "STUB: not implemented"; return false }

// Count returns the number of elements in a collection.
func (q Query) Count() int { _ = "STUB: not implemented"; return 0 }

// CountWith returns a number that represents how many elements in the specified
// collection satisfy a condition.
func (q Query) CountWith(predicate func(any) bool) int { _ = "STUB: not implemented"; return 0 }

// CountWithT is the typed version of CountWith.
//
//   - predicateFn is of type "func(TSource) bool"
//
// NOTE: CountWith has better performance than CountWithT.
func (q Query) CountWithT(predicateFn any) int { _ = "STUB: not implemented"; return 0 }

// First returns the first element of a collection.
func (q Query) First() any { _ = "STUB: not implemented"; return *new(any) }

// FirstWith returns the first element of a collection that satisfies a
// specified condition.
func (q Query) FirstWith(predicate func(any) bool) any { _ = "STUB: not implemented"; return *new(any) }

// FirstWithT is the typed version of FirstWith.
//
//   - predicateFn is of type "func(TSource) bool"
//
// NOTE: FirstWith has better performance than FirstWithT.
func (q Query) FirstWithT(predicateFn any) any { _ = "STUB: not implemented"; return *new(any) }

// ForEach performs the specified action on each element of a collection.
func (q Query) ForEach(action func(any)) { _ = "STUB: not implemented"; return }

// ForEachT is the typed version of ForEach.
//
//   - actionFn is of type "func(TSource)"
//
// NOTE: ForEach has better performance than ForEachT.
func (q Query) ForEachT(actionFn any) { _ = "STUB: not implemented"; return }

// ForEachIndexed performs the specified action on each element of a collection.
//
// The first argument to action represents the zero-based index of that
// element in the source collection. This can be useful if the elements are in a
// known order and you want to do something with an element at a particular
// index, for example. It can also be useful if you want to retrieve the index
// of one or more elements. The second argument to action represents the
// element to process.
func (q Query) ForEachIndexed(action func(int, any)) { _ = "STUB: not implemented"; return }

// ForEachIndexedT is the typed version of ForEachIndexed.
//
//   - actionFn is of type "func(int, TSource)"
//
// NOTE: ForEachIndexed has better performance than ForEachIndexedT.
func (q Query) ForEachIndexedT(actionFn any) { _ = "STUB: not implemented"; return }

// Last returns the last element of a collection.
func (q Query) Last() (r any) { _ = "STUB: not implemented"; return *new(any) }

// LastWith returns the last element of a collection that satisfies a specified
// condition.
func (q Query) LastWith(predicate func(any) bool) (r any) {
	_ = "STUB: not implemented"
	return *new(any)
}

// LastWithT is the typed version of LastWith.
//
//   - predicateFn is of type "func(TSource) bool"
//
// NOTE: LastWith has better performance than LastWithT.
func (q Query) LastWithT(predicateFn any) any { _ = "STUB: not implemented"; return *new(any) }

// Max returns the maximum value in a collection of values.
func (q Query) Max() any { _ = "STUB: not implemented"; return *new(any) }

// Min returns the minimum value in a collection of values.
func (q Query) Min() any { _ = "STUB: not implemented"; return *new(any) }

// Results collects all items from a query into a slice.
func (q Query) Results() []any { _ = "STUB: not implemented"; return nil }

// SequenceEqual determines whether two collections are equal.
func (q Query) SequenceEqual(q2 Query) bool { _ = "STUB: not implemented"; return false }

// Single returns the only element of a collection, and nil if there is not
// exactly one element in the collection.
func (q Query) Single() (r any) { _ = "STUB: not implemented"; return *new(any) }

// SingleWith returns the only element of a collection that satisfies a
// specified condition, and nil if more than one such element exists.
func (q Query) SingleWith(predicate func(any) bool) (r any) {
	_ = "STUB: not implemented"
	return *new(any)
}

// SingleWithT is the typed version of SingleWith.
//
//   - predicateFn is of type "func(TSource) bool"
//
// NOTE: SingleWith has better performance than SingleWithT.
func (q Query) SingleWithT(predicateFn any) any { _ = "STUB: not implemented"; return *new(any) }

// SumInts computes the sum of a collection of numeric values.
//
// Values can be of any integer type: int, int8, int16, int32, int64. The result
// is int64. Method returns zero if the collection contains no elements.
func (q Query) SumInts() (r int64) { _ = "STUB: not implemented"; return 0 }

// SumUInts computes the sum of a collection of numeric values.
//
// Values can be of any unsigned integer type: uint, uint8, uint16, uint32,
// uint64. The result is uint64. Method returns zero if the collection contains no
// elements.
func (q Query) SumUInts() (r uint64) { _ = "STUB: not implemented"; return 0 }

// SumFloats computes the sum of a collection of numeric values.
//
// Values can be of any float type: float32 or float64. The result is float64.
// Method returns zero if the collection contains no elements.
func (q Query) SumFloats() (r float64) { _ = "STUB: not implemented"; return 0 }

// ToChannel iterates over a collection and outputs each element to a channel,
// then closes it.
func (q Query) ToChannel(result chan<- any) { _ = "STUB: not implemented"; return }

// ToChannelT is the typed version of ToChannel.
//
//   - result is of type "chan TSource"
//
// NOTE: ToChannel has better performance than ToChannelT.
func (q Query) ToChannelT(result any) { _ = "STUB: not implemented"; return }

// ToMap iterates over a collection and populates a result map with elements.
// Collection elements have to be of KeyValue type to use this method. To
// populate a map with elements of different types, use the ToMapBy method. ToMap
// doesn't empty the result map before populating it.
func (q Query) ToMap(result any) { _ = "STUB: not implemented"; return }

// ToMapBy iterates over a collection and populates the result map with
// elements. Functions keySelector and valueSelector are executed for each
// element of the collection to generate key and value for the map. Generated
// key and value types must be assignable to the map's key and value types.
// ToMapBy doesn't empty the result map before populating it.
func (q Query) ToMapBy(result any,
	keySelector func(any) any,
	valueSelector func(any) any) {
	_ = "STUB: not implemented"
	return
}

// ToMapByT is the typed version of ToMapBy.
//
//   - keySelectorFn is of type "func(TSource)TKey"
//   - valueSelectorFn is of type "func(TSource)TValue"
//
// NOTE: ToMapBy has better performance than ToMapByT.
func (q Query) ToMapByT(result any,
	keySelectorFn any, valueSelectorFn any) {
	_ = "STUB: not implemented"
	return
}

// ToSlice iterates over a collection and saves the results in the slice pointed
// by v. It overwrites the existing slice, starting from index 0.
//
// If the slice pointed by v has sufficient capacity, v will be pointed to a
// resliced slice. If it does not, a new underlying array will be allocated and
// v will point to it.
//
// Note: Starting with go-linq v4, ToSlice panics if v is not a pointer to a slice.
// If the query type is not assignable to the slice element type, ToSlice will
// attempt to convert the query elements to the slice element type.
func (q Query) ToSlice(v any) { _ = "STUB: not implemented"; return }

// Reset length to 0 but keep capacity (like s = s[:0])
// This preserves any existing capacity for reuse.

// Ensure type compatibility with the slice element type.

// Point v to the final slice (which may have a new backing array).
