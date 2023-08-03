package Value

import (
	"fmt"
)

type value struct {
	Data float64
	Prev Pair[*value, *value]
}

func CreateValue(data float64) value {
	return value{data, Pair[*value, *value]{nil, nil}}
}

func (v value) GetData() {
	fmt.Println("Data: ", v.Data)
}

// Add returns the sum of two Floats Values
func (t value) Add(other value) value {
	return value{t.Data + other.Data, Pair[*value, *value]{&t, &other}}
}

// Mul returns the multiplication of two Floats Values
func (t value) Mul(other value) value {
	return value{t.Data * other.Data, Pair[*value, *value]{&t, &other}}
}

type Pair[T, U any] struct {
	First  T
	Second U
}
