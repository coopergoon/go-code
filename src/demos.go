package main

import (
	"./demos"
	"log"
)

var renlen, renwith float64 = 1, 2

func init()  {
	//println("-------1111112222")
	if renlen > 10 {
		log.Fatal("LEN IS SHORT")
	}
	if renwith > 2 {
		log.Fatal("with11111")
	}
	
}

func main()  {
	//fmt.Println("----2")
	//fmt.Println("___3", demos.Demo1(renwith, renlen))
	//demos.Demo2()
	//demos.Demo3()
	//demos.Demo4()
	//demos.Demo5()
	//demos.Demo6x()
	//demos.Demo7()
	//demos.Demo8()
	//demos.Demo9()
	//demos.Demo10()
	//demos.Demo11()
	//demos.Demo12()
	//demos.Demo13()
	//demos.Demo14()
	//demos.Demo15()
	//demos.Demo16()
	//demos.Demo17()
	//demos.Demo18()
	//demos.Demo19()
	//demos.TypeDemo()
	//demos.FuncDemo()
	//demos.Func3()
	//demos.FuncDemo1()
	//demos.FuncDemo5()
	demos.InterfaceDemo()
}



