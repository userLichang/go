package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数的程序
1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2.开启24个goroutine从jobChan中取出随机数计算各个数的和，将结果发送到resultChan
3.主goroutine从resultChan取出结果并打印到终端输出
 */

type  job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var  wg = sync.WaitGroup{}


func chengcheng(cc chan<- *job)  {
	defer wg.Done()
	//循环生成int64类型的随机数，发送到jobChan
	for  {
		x := rand.Int63()
		newJob := &job{
			value : x,
		}
		cc <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}
func yushi(cc <-chan *job, resultChan chan<- *result)  {
	defer wg.Done()
	//从jobChan中取出随机数计算哥位数的和，将结果发送到resulChan
	for  {
		job := <-cc
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult

	}
}
func main()  {
	wg.Add(1)
	go chengcheng(jobChan)
	for i := 0; i < 24; i++ {
		go yushi(jobChan, resultChan)
	}
	//主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}