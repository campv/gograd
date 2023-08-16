package Value

import (
	"fmt"
)

type Value struct {
	Data float64
	Prev Tuple[*Value, *Value]
	Op   string
}

func CreateValue(data float64) Value {
	return Value{data, Tuple[*Value, *Value]{nil, nil}, ""}
}

func (v Value) GetData() {
	fmt.Println("Data: ", v.Data)
}

// Add returns the sum of two Floats Values
func (t Value) Add(other Value) Value {
	return Value{t.Data + other.Data, Tuple[*Value, *Value]{&t, &other}, "+"}
}

// Mul returns the multiplication of two Floats Values
func (t Value) Mul(other Value) Value {
	return Value{t.Data * other.Data, Tuple[*Value, *Value]{&t, &other}, "*"}
}

// Tuple is a generic type that represents a pair of Values.
// The first Value is of type T, and the second Value is of type U.
type Tuple[T, U any] struct {
	First  T
	Second U
}

// TODO implement the Value struct using generic type for numbers
// Number is a type that represents a number standard type.
type Number interface {
	int64 | float64 | complex128 |
		int32 | float32 | complex64 |
		int16 | uint64 |
		int8 | uint32 |
		uint16 | uint8 |
		int | uint
}
