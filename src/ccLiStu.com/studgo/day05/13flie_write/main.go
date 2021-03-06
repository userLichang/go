package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容
func wirteDemo1() {
	fileObj, err := os.OpenFile("./xx.go", os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("chengcheng ai shishi!"))
	//writeString
	fileObj.WriteString("诗诗是谁啊！")
	//关闭文件
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.go", os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello shishi\n")  //写到缓存
	wr.Flush()       	//将缓存中的内容写入文件

}

func writeDemo3()  {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err", err)
		return
	}
}
func main()  {
	//writeDemo2()
	writeDemo3()
}