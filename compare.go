package linq

type comparer func(any, any) int

// Comparable is an interface that has to be implemented by a custom type
// to work with linq.
//
// Example:
//
//	func (f foo) CompareTo(c Comparable) int {
//		a, b := f.f1, c.(foo).f1
//
//		if a < b {
//			return -1
//		} else if a > b {
//			return 1
//		}
//
//		return 0
//	}
type Comparable interface {
	CompareTo(Comparable) int
}

func getComparer(data any) comparer { _ = "STUB: not implemented"; return *new(comparer) }
