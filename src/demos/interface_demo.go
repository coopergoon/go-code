package demos

import "fmt"

func InterfaceDemo()  {
	//同一个包下调用另一个文件中的结构体
	em1 := ep{

		name:"name1",
		age: 10002,

	}
	fmt.Println(em1)
}