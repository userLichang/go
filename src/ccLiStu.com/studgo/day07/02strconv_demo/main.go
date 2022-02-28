package main

import (
	"fmt"
	"strconv"
)

//strconv

func main()  {
	//从字符串中解析出对应的数据
	str := "1000"
	//ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, int(ret1))

	 //Atoi 字符串转换成int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	//从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)
	//从字符串中解析出浮点类型
	floatStr := "1.24"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)
	i := 97
	//ret2 := string(i)  //"a"

	ret2 := fmt.Sprintf("%d", i) //97
	fmt.Printf("%#v", ret2)
}
