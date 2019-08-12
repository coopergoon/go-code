package demo1

import (
	"fmt"
	"math"
)

//能不能不共享同一个切片

func init()  {
	fmt.Println("12345上山打老虎")
}

func Demo1(a, b float64) (c float64) {
	c = math.Sqrt(a * b)
	return c
}


func Demo2(){
	if a1 := 10; a1%2 > 3{
		fmt.Println("q1")
	} else {
		fmt.Println("q2")
	}
}

func Demo3()  {
	for i:=1;i<10;{
		if i == 3 {
			fmt.Println(i)
		} else if  i == 6 {
			fmt.Println("---23ffd")
			break

		}
		fmt.Println(i)
		i+=1
	}
}

func Demo4()  {
	a := [...]int{1,2,34}
	fmt.Println(a)
	fmt.Println(a[0])
	a[0] = 12
	b := a
	b[0] = 1222
	fmt.Println(a)
	fmt.Println(b)
}

func Demo5()  {
	a := [...]int{11,22,33,4,45,6,7,8}
	for i,v := range a {
		fmt.Printf("%d is %d\n", i,v )
	}
	fmt.Println("____12_____")
	for i:=0; i<len(a); i++ {
		fmt.Println(a[i])
	}
}

func Demo6(d [1][2]string)  {
	a := [2][3]string{
		{"a", "b"},
		{"a1", "b1"},
	}

	for i1, v1 := range a{
		fmt.Println(v1)
		for i2, v2 := range v1 {
			fmt.Println(i1, i2, v2)
		}
	}

	fmt.Println("--------q2")
	fmt.Println(d)
}

func Demo6x(){
	d := [1][2]string{
		{"1", "2"},
	}
	Demo6(d)

}

func Demo7()  {
	c := []int{1,2,3}
	fmt.Println(c[1:])
}

func Demo8()  {
	a := [...]int{1,2,3,4,5,6}
	b := a[2:]
	c := a[2:]
	b[1] = 23344
	for i, v := range b {
		fmt.Println(i, v)
	}

	c[1] = 12233
	fmt.Println("=======123")
	for i, v := range c {
		fmt.Println(i, v)
	}
	fmt.Println("=======122222221-----3")
	for i, v := range b {
		fmt.Println(i, v)
	}
	fmt.Println(len(c), cap(c), c)


}



func Demo9(){
	i := make([]int, 2, 3)
	fmt.Println(cap(i), len(i))
	fmt.Println(i)
	i[2] = 1
	//i[3] = 1
	fmt.Println(i)

}


func Demo10(){
	a := []int{1,2,3}
	fmt.Println(a, len(a), cap(a))
	b := append(a, 4)
	fmt.Println(b, len(b), cap(b))
}


func Demo11(){
	a := []int{1,2,3,4}
	b := []int{5,6,7,8,9}
	c := append(a, b...)
	fmt.Println(c)

}

func one(nu []int)  {
	for i := range nu{
		nu[i] -= 2
	}
}

func Demo12()  {
	a := []int{10, 12, 14}
	one(a)
	fmt.Println(a)
}

func country() []int {
	a := []int{1,2,3,4,5,6}
	a1 := a[:len(a) - 2]
	a2 := make([]int, len(a1))
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	copy(a2, a1)
	return a2
}

func Demo13(){
	d := country()
	fmt.Println(d)

}
func find(num string, nums ...string){
	fmt.Println(num)
	fmt.Println(nums)
	f := false
	for i, v := range nums{
		if v == num{
			fmt.Println(i, v, num)
			fmt.Println("-------q1")
			f = true
		}
	if !f{
		fmt.Println("--------12")
	}

	}
}

func Demo14()  {
	nums := []string{"1", "2", "3"}
	find("11", "11","22","33", "44")
	find("11", nums...)
}

func demo15(s ...int)  {
	s[0] = 111111
	fmt.Println(s)
	//s = append(s, 22222)
	s1 := append(s, 22222)
	//fmt.Println(s)
	fmt.Println(s1)
}

func Demo15(){
	a := []int{1,2,3,4}
	demo15(a...)
}

func Demo16()  {
	d := make(map[string]int)
	fmt.Println(d)
	d["1q"] = 1
	d["2q"] = 11111111111111111
	d["3q"] = 1
	fmt.Println(d)
	var s map[string]int
	fmt.Println(s)
	if s == nil{
		fmt.Println("___21")
	} else {
		fmt.Println("___22221")
	}

	f := map[string]int{
		"11":22,
		"22":22,
	}
	fmt.Println(f)
	f["222"] = 1244
	fmt.Println(d["2111"])
	_, m := f["2123"]
	if m == false {
		fmt.Println(m)
	}


	for i, v := range d {
		fmt.Println(i, v)
	}
	delete(d, "1q")
	fmt.Println(d)

	aa := "sdfghh"
	fmt.Println(len(aa))

}



func Demo17(){
	a := "dsds"
	//var b *string = &a
	b := &a
	var c *int
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(*b)
}

type aa struct {
	name, local string
	age, time int
}

func Demo18()  {
	a := aa{
		name: "dsds",
		local: "dw",
		age: 12,
		time: 3232,
	}

	fmt.Println(a)

	b := aa{"ds", "ewe", 12,34}

	fmt.Println(b)
	fmt.Println(a)


	ds := struct {
		name string
		age int
	}{
		name:"dd",
		age:122,
	}
	fmt.Println(ds)

}

type  em struct {
	name, loacal string
	age int
}

func Demo19(){
	var a em
	fmt.Println(a)

}