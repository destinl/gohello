package gohello

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Concurrent() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for a := range abc {
			fmt.Println("a: ", a)
		}
		fmt.Println("end")
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}

func Defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
	// 输出：
	// 打印后
	// 打印中
	// 打印前
	// panic: 触发异常
}

type student struct {
	Name string
	Age  int
}

func Pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(m)
}

func Func_I() {
	runtime.GOMAXPROCS(1) //只有1个操作系统线程可供用户的Go代码
	//从runtime的源码可以看到，当创建一个G时，会优先放入到下一个调度的runnext字段上作为下一次优先调度的G。
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
