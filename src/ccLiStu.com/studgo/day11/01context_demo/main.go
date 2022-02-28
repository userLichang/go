package main

import (
	"fmt"
	"sync"
	"time"
)

//为什么需要context
var wg sync.WaitGroup
var notify bool

func f()  {
	defer wg.Done()
	for {
		fmt.Println("诗妹")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main()  {

	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	//如何通知goroutine退出
	notify = true
	wg.Wait()

}