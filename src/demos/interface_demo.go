package demos

import "fmt"

type myString string

type vom interface {
	FindMe() []rune
}

func (m myString)FindMe() []rune {
	var vowels []rune
	for _, rune := range m{
		//fmt.Println("0", s)
		//fmt.Println("1", rune)
		if rune == 'a' || rune == 'e'{
			vowels = append(vowels, rune)
		}

	}
	return vowels
}

func InterfaceDemo()  {
	//同一个包下调用另一个文件中的结构体
	//em1 := ep{
	//
	//	name:"name1",
	//	age: 10002,
	//
	//}
	//fmt.Println(em1)

	name := myString("eweweweavbr")
	var v vom
	v = name
	fmt.Printf(v.FindMe())

}

/*
	接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。
*/


