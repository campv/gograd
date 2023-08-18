package main

import (
	"fmt"

	Value "github.com/campv/gograd/value"
)

func main() {
	fmt.Println("gograd started")

	/*
		a := 2.0
		b := -3.0
		c := 10.0
		d := a*b + c

		//print d value
		fmt.Println(d)

		h := 0.0001

		d1 := d
		c += h
		d2 := a*b + c

		fmt.Println("d1: ", d1)
		fmt.Println("d2: ", d2)
		fmt.Println("slope: ", (d2-d1)/h)
	*/
	/*
		a := Value.CreateValue(2.0)
		b := Value.CreateValue(-3.0)
		c := Value.CreateValue(10.0)

		a.Label = "a"
		b.Label = "b"
		c.Label = "c"
		a.GetData()
		b.GetData()
		c.GetData()

		e := a.Add(b)
		e.Label = "e"
		d := e.Add(c)
		d.Label = "d"
		f := Value.CreateValue(-2.0)
		f.Label = "f"
		L := d.Mul(f)
		L.GetData()
		fmt.Println(L.Prev.First.Data)
		fmt.Println(L.Prev.Second.Data)
		fmt.Println(L.Op)
	*/
	// activation neuron
	// inputs
	x1 := Value.CreateValue(2.0)
	x2 := Value.CreateValue(0.0)
	// weights
	w1 := Value.CreateValue(-3.0)
	w2 := Value.CreateValue(1.0)
	// bias
	b := Value.CreateValue(6.88137335870195432)
	x1.Label = "x1"
	x2.Label = "x2"
	w1.Label = "w1"
	w2.Label = "w2"
	b.Label = "b"

	// w1*x1 + w2*x2 + b
	z := w1.Mul(x1).Add(w2.Mul(x2)).Add(b)
	z.Label = "z"

	// activation function
	a := z.Tanh()
	a.Label = "a"

	a.ShowValue()
	fmt.Println(a.Prev.First.Data)

	fmt.Println(a.Op)
	fmt.Println(a.Label)
	// initialize gradient
	a.Grad = 1.0
	fmt.Println(a.Grad)
	fmt.Println(a.Prev.First.Grad)
	// backpropagate
	fmt.Println("Backpropagate")
	a.Backward()
	fmt.Println(a.Prev.First.Grad)
}
