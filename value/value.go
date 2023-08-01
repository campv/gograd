package Value

import (
	"fmt"
)

type Value interface {
	GetData()
	Add(Value) Value
}

type Float struct {
	Data float64
}

type Int struct {
	Data int64
}

func (v Float) GetData() {
	fmt.Println("Data: ", v.Data)
}

func (v Int) GetData() {
	fmt.Println("Data: ", v.Data)
}

// Add returns the sum of two Floats Values
func (t Float) Add(other Float) Float {
	return Float{t.Data + other.Data}
}

// Mul returns the multiplication of two Floats Values
func (t Float) Mul(other Float) Float {
	return Float{t.Data * other.Data}
}
