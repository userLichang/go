package main

import (
	"fmt"
	"runtime"
	"sync"
)

//gomaxprocs
var wg sync.WaitGroup

func a()  {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b()  {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main()  {
	runtime.GOMAXPROCS(6)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}