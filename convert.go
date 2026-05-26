package linq

type intConverter func(any) int64

func getIntConverter(data any) intConverter { _ = "STUB: not implemented"; return *new(intConverter) }

type uintConverter func(any) uint64

func getUIntConverter(data any) uintConverter {
	_ = "STUB: not implemented"
	return *new(uintConverter)
}

type floatConverter func(any) float64

func getFloatConverter(data any) floatConverter {
	_ = "STUB: not implemented"
	return *new(floatConverter)
}
