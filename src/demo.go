package main

import "fmt"
import "./demos"

func main()  {

	a, b := 10, 20
	c ,d := test(a, b)
	fmt.Println("ddsdsd", c, d)

	fmt.Println("___________")
	c1, c2 := test1(a, b)
	fmt.Println("a", c1, c2)

	var a1, b2 float64 = 4, 5
	c3 := demos.Demo1(a1, b2)
	fmt.Println(c3)
}


func test(a int, b int) (c, d int) {
	c = a * b
	d = a * b * 1000
	return
}


func test1(a, b int) (int, int) {
	c := a * b
	d := a * b * 1000
	return c, d
}