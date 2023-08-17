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

}
