package main

func main() {
	// needs to be full struct
	test1 := Cos{1, 2, 3, 4.0, 5.0, 6.0}
	test2 := Cos{1, 2, 3, 4.0, 5.0, 6.0}
	test3 := test1.Add(test2)
	println(test3.x)

	x := 1
	y := 2
	z := 3
	a := 4.0

	// optional using pointers, can be partial struct (only the fields that are needed) and can be in any order
	// undefined fields are nil
	test1optional := Cos2{
		x: &x,
		y: &y,
		z: &z,
		a: &a,
	}

	test2optional := Cos2{
		x: &x,
		c: &a,
	}
	println(*test1optional.x + *test2optional.x)
}

type Cos struct {
	x, y, z int
	a, b, c float64
}

// optional using pointers
type Cos2 struct {
	x, y, z *int
	a, b, c *float64
}

func (p Cos) Add(q Cos) Cos {
	return Cos{p.x + q.x, p.y + q.y, p.z + q.z, p.a + q.a, p.b + q.b, p.c + q.c}
}
