package main

import "fmt"

func main() {
	var a int = 12
	var b float32 = 4.21
	const pi float64 = 3.1415139457

	var (
		d string = "haha"
		e string = "hihi"
		f string = "hoho"
	)

	x, y, z := 1, 2, 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(d, e, f)
	fmt.Println(x, y, z)

}
