package Value

import (
	"fmt"
)

type value struct {
	Data float64
	Prev Tuple[*value, *value]
	Op   string
}

func CreateValue(data float64) value {
	return value{data, Tuple[*value, *value]{nil, nil}, ""}
}

func (v value) GetData() {
	fmt.Println("Data: ", v.Data)
}

// Add returns the sum of two Floats Values
func (t value) Add(other value) value {
	return value{t.Data + other.Data, Tuple[*value, *value]{&t, &other}, "+"}
}

// Mul returns the multiplication of two Floats Values
func (t value) Mul(other value) value {
	return value{t.Data * other.Data, Tuple[*value, *value]{&t, &other}, "*"}
}

// Tuple is a generic type that represents a pair of values.
// The first value is of type T, and the second value is of type U.
type Tuple[T, U any] struct {
	First  T
	Second U
}

// TODO implement the value struct using generic type for numbers
// Number is a type that represents a number standard type.
type Number interface {
	int64 | float64 | complex128 |
		int32 | float32 | complex64 |
		int16 | uint64 |
		int8 | uint32 |
		uint16 | uint8 |
		int | uint
}
