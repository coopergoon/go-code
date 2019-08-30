package demos

import "fmt"

/*
go的函数调用包括两种：1、有接收器的，我们称之为方法；2、无接收器的，我们称之为函数；
方法又包含两种：
1、指针方法：指针做为接收器；
2、值方法：值做为接收器；

指针方法可以改变对象的field，相当于C里的指针；
值方法不会改变对象的field
一个对象包含的所有值方法称为值方法集；
一个对象包含的所有指针方法，称为指针方法集；
go的参数传递总是值传递，只不过如果传递的是指针，那么可以改变指针所指对象；

关于接收器有两个非常重要的规则，必须要牢记，对于设计接口非常有必要知道：

值只包含值方法集，不包含指针方法集；
指针包含值方法集，也包含值方法集（全部都有就对了）
所以有时我们需要将一个对象赋值给接口时，如果传值，但是值只实现了接口的值方法，而没实现指针方法，就会报运行时错误；
而如果传递指针，肯定没有问题；
*/

type Ts struct {
	name string 
	age int
}

type Tf struct {
	time string
	number int
}

func (t Tf) f1 ()  {
	fmt.Println(t.time)
	fmt.Println(t.number)
}


//接收器类型为Ts  这是个方法 没有接收器的叫做函数
func (t Ts) f1 ()  {
	fmt.Println(t.age)
	fmt.Println(t.name)
}

func FuncDemo()  {
	es1 := Ts{
		name: "chen",
		age:  0,
	}

	es2 := Tf{
		time: "2019",
		number: 999,
	}

	es1.f1()  // 调用es1类型的 f1 方法
	es2.f1()   // 调用es2类型的 f1 方法
}

/*
既然我们可以使用函数写出相同的程序，那么为什么我们需要方法？这有着几个原因，让我们一个个的看看。
Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
*/



//值接收器和指针接收器 

type demo0 struct {
	name string 
	age int
}

//值接受方法  不能改变   name不会被赋值为新的值
func (d demo0)func1 (name string)  {
	d.name = name
}

//指针接受方法  可以改变内部的变量   age会被赋值为新的值
func (d *demo0) func2 (age int)  {
	d.age = age
}

func Func3()  {
	d1 := demo0{
		name: "chen09876",
		age:  0,
	}

	fmt.Println(d1.name)
	d1.func1("chenoppppp")   // demo0 结构体的 func1方法
	fmt.Println(d1.name)

	fmt.Println(d1.age)
	d1.func2(1000000)
	fmt.Println(d1.age)


}

type rect struct {
	a, b string
}

//这是一个普通的函数
func b1(r rect)  {
	fmt.Println(r.a)
	fmt.Println(r.b)
}


//这是一个值接收器
func (r rect) b1 ()  {
	fmt.Println(r.a)
	fmt.Println("________________")
	fmt.Println(r.b)
}

func FuncDemo1()  {
	p := rect{
		a: "qp",
		b: "qp2",
	}

	//函数
	b1(p)   // 常见的函数调用
	p.b1()   // rect 结构体调用 b1方法  值接收器

	p1 := &p // 创建 p结构体的指针 p1
	p1.b1()   // 使用p1 指针调用值接收器

	/*
	函数的参数是不能接受指针的
	b1(p1) 所以这样的写法是错误的
	*/

}



type per1 struct {
	name string
	age int
}

//函数  接受指针参数
func mw1(p *per1)  {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

//指针接收器方法
func (p per1) mw2 ()  {
	fmt.Println(p.name)
	fmt.Println("---------")
	fmt.Println(p.age)
}




func FuncDemo5()  {
	p5 := per1{
		name: "mbmbmbm",
		age: 13312321,
	}
	p6 := &p5  // 创建指针
	mw1(p6)  // 函数接受指针参数

	//p5结构体调用指针接收器方法
	p5.mw2()

	p6.mw2()   // 指针调用指针接收器方法

	//不可以使用将值
	//p6.mw1()  这是错误的


	//在非结构体上的方法
	//这部分没明白咋回事

	num1 := myint(10)
	num2 := myint(10000000000)
	num1.p(num2)


}

type myint int

func (d myint)p(a myint) {
	//return a + d
	fmt.Println("------cc" , a + d)
}

