package main

import "fmt"

//给自定义类型加方法
//不能给别的包里面的类型添加方法，只能给自己的包里
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}
func main()  {
	m := myInt(100) //int8(10)
	m.hello()
}