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

	a := Value.Float{Data: 2.0}
	b := Value.Float{Data: -3.0}
	c := Value.Float{Data: 10.0}

	a.GetData()
	b.GetData()
	c.GetData()

	d := a.Mul(b).Add(c)
	d.GetData()

}
