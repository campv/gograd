package Value

import (
	"fmt"
	"math"
)

type Value struct {
	Data     float64
	Grad     float64
	backward func(out *Value)
	Prev     Tuple[*Value, *Value]
	Op       string
	Label    string
}

func CreateValue(data float64) Value {
	return Value{data, 0.0, func(out *Value) {}, Tuple[*Value, *Value]{nil, nil}, "", ""}
}

func (v Value) ShowValue() {
	fmt.Printf("Data: %f\nGradient: %f\nOp: %s", v.Data, v.Grad, v.Op)
}

func (v Value) Backward() {
	v.backward(&v)
}

// Add returns the sum of two Floats Values
func (t Value) Add(other Value) Value {
	out := Value{t.Data + other.Data, 0.0, nil, Tuple[*Value, *Value]{&t, &other}, "+", t.Label + other.Label}
	backward := func(out *Value) {
		t.Grad = 1.0 * out.Grad
		other.Grad = 1.0 * out.Grad
	}
	out.backward = backward
	return out
}

// Mul returns the multiplication of two Floats Values
func (t Value) Mul(other Value) Value {
	out := Value{t.Data * other.Data, 0.0, nil, Tuple[*Value, *Value]{&t, &other}, "*", t.Label + other.Label}
	backward := func(out *Value) {
		t.Grad = other.Data * out.Grad
		other.Grad = t.Data * out.Grad
	}
	out.backward = backward
	return out
}

// Tanh returns the hyperbolic tangent of the Float Value
func (t Value) Tanh() Value {
	out := Value{math.Tanh(t.Data), 0.0, nil, Tuple[*Value, *Value]{&t, nil}, "tanh", "tanh(" + t.Label + ")"}
	backward := func(out *Value) {
		t.Grad = (1.0 - math.Pow(out.Data, 2.0)) * out.Grad
	}
	out.backward = backward
	return out
}

// Sigmoid returns the sigmoid of the Float Value
func (t Value) Sigmoid() Value {
	return Value{t.Data, 0.0, nil, Tuple[*Value, *Value]{&t, nil}, "sigmoid", "sigmoid(" + t.Label + ")"}
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
