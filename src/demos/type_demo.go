package demos

import "fmt"
//结构图是用户自定义的结构 里面可以包含很多个集合和类型
//结构体的命名可以是小写字母

type ep struct {
	name string
	home string
	age int
}

func TypeDemo()  {

	//这样的方式只需要填写自己想填写的参数
	em1 := ep{
		name: "chen122345",
		age: 122,
		//如果没有设置某个字段的值，string会被赋值为空 int会是0
	}

	//这样的方式必须填写每一个参数
	em2 := ep{"chen", "dsde", 232}
	fmt.Println(em1)
	fmt.Println(em2)



	//创建匿名结构体
	//就是省略了结构体的名称，并且紧接着定义结构体的内部内容

	em3 := struct {
		names, homes string
		ages int
	}{
		names: "names1",
		homes: "homes2",
		ages: 123455,

	}

	fmt.Println("em3: ", em3)

	var em4 ep
	fmt.Println(em4)
	//输出的是 {  0}   因为字符串是空值，  int是问0



	//访问结构体 用操作符 . 来访问
	n := em3.names
	a := em3.ages
	fmt.Println("get name: " + n)
	fmt.Println("get name: ", a)


	//创建指向结构体的指针
	type bm struct {
		name, time string
		age int
	}


	//em5 是一个指向bm结构体的指针
	em5 := &bm{
		name: "bob",
		time: "2019",
		age: 010001,
	}
	//（*em5）是指针  也可以直接使用em5
	a1 := (*em5).name
	a12 := em5.name
	a2 := (*em5).time
	a3 := (*em5).age

	fmt.Println("a1: " + a1)
	fmt.Println("a1: " + a12)
	fmt.Println("a2: " + a2)
	fmt.Println("a3: ", a3)


	//嵌套结构体
	type m1 struct {
		name string
		age int
	}

	type m2 struct {
		home string
		m3 m1
	}

	var m m2
	m.m3 = m1{
		name: "chen",
		age:  1000,
	}

	m.home = "ap"

	fmt.Println(m)






}