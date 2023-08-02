package Value

import (
	"fmt"
)

type Value struct {
	Data int64
}

func (v Value) GetData() {
	fmt.Println("Data: ", v.Data)
}

// Add returns the sum of two Floats Values
func (t Value) Add(other Value) Value {
	return Value{t.Data + other.Data}
}

// Mul returns the multiplication of two Floats Values
func (t Value) Mul(other Value) Value {
	return Value{t.Data * other.Data}
}
