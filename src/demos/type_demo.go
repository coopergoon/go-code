package demos

import "fmt"
//结构图是用户自定义的结构 里面可以包含很多个集合和类型
//结构体的命名可以是小写字母

type ep struct {
	name, home string
	age int
}

func TypeDemo()  {

	//这样的方式只需要填写自己想填写的参数
	em1 := ep{
		name: "chen",
		age: 122,
	}

	//这样的方式必须填写每一个参数
	em2 := ep{"chen", "dsde", 232}
	fmt.Println(em1)
	fmt.Println(em2)
	
}