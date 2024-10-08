package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gohello/gohello"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var tr *http.Transport

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
		// 下面的代码被干掉了
		//Dial: func(netw, addr string) (net.Conn, error) {
		//	conn, err := net.DialTimeout(netw, addr, time.Second*2) //设置建立连接超时
		//	if err != nil {
		//		return nil, err
		//	}
		//	err = conn.SetDeadline(time.Now().Add(time.Second * 3)) //设置发送接受数据超时
		//	if err != nil {
		//		return nil, err
		//	}
		//	return conn, nil
		//},
	}
}

func Get(url string) ([]byte, error) {
	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	req, _ := http.NewRequest("Get", url, body)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second, // 超时加在这里，是每次调用的超时
	}
	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

func goroutine_channel() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

func goroutine_channel_2() {
	random := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			num, ok := <-random
			if ok {
				fmt.Println(num)
			} else {
				done <- true
			}
		}
	}()

	go func() {
		defer close(random)
		for i := 0; i < 5; i++ {
			random <- rand.Intn(5)
		}
	}()

	<-done
	close(done)
}

func SearchTarget(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "Task cancelded! \n")
			return
		default:
		}
		// 模拟一个耗时查找，这里只是比对值，真实开发中可以是其他操作
		fmt.Fprintf(os.Stdout, "v: %d \n", v)
		time.Sleep(time.Millisecond * 1500)
		if target == v {
			resultChan <- true
			return
		}
	}

}

func multi_goroutine() {
	timer := time.NewTimer(time.Second * 5)
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dataLen := len(data)
	size := 3
	target := 345
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)
	for i := 0; i < dataLen; i += size {
		end := i + size
		if end >= dataLen {
			end = dataLen - 1
		}
		go SearchTarget(ctx, data[i:end], target, resultChan)
	}
	select {
	case <-timer.C:
		fmt.Fprintln(os.Stderr, "Timeout! Not Found")
		cancel()
	case <-resultChan:
		fmt.Fprintf(os.Stdout, "Found it!\n")
		cancel()
	}

	time.Sleep(time.Second * 2)
}

func closures() {
	funcs := make([]func(), 3)
	for i := 0; i < 3; i++ {
		// funcs[i] = func() {
		// 	fmt.Println(i)
		// 	// 闭包是一个函数值，它可以访问外部函数的变量，即使外部函数已经返回了。
		// 	// 这里的i是闭包中捕获的变量，而不是函数参数
		// 	// 0, 1, 2 (3 3 3就是使用闭包可能的错误情况)
		// }
		// 创建局部变量j，在闭包中引用这个局部变量，这样，每个闭包就会捕获不同的变量值
		j := i
		funcs[i] = func() {
			fmt.Println(j)
		}
	}
	for _, f := range funcs {
		f()
	}
}

func createSlice() []int {
	//这里的切片会逃逸到堆上，因为它是通过值返回的（它的生命周期超出了函数的执行范围，比函数长）
	return []int{1, 2, 3}
}

func createSlice2() *[]int {
	//这里的切片不会逃逸到堆上，因为它是通过指针返回的
	s := []int{1, 2, 3}
	return &s
}

func main() {
	m := map[string]gohello.Student{"people": {"zhoujielun"}}
	student := m["people"]   // 获取副本
	student.Name = "wuyanzu" // 修改副本的字段
	m["people"] = student    // 将修改后的值放回 map
	fmt.Println(m)

	// if gohello.Live() == nil {
	// 	fmt.Println("AAAAAAA")
	// } else {
	// 	fmt.Println("BBBBBBB")
	// }

	// t := gohello.Teacher{}
	// t.ShowA()

	// gohello.Func_I()

	// gohello.Pase_student()

	// gohello.Defer_call()

	// gohello.Concurrent()

	// goadmin.Goadmin()

	// crawler.GetPicture()

	// utils.GenerateUUID()

	// utils.Cron()

	// utils.TcpScanner()

	// websocket.Websocket_main()

	// failed:upgrade error: websocket: the client is not using the websocket protocol: 'upgrade' token not found in 'Connection' header
	// gohello.StartWebSocketServer()

	// s := createSlice()
	// fmt.Println(s)
	// s2 := createSlice2()
	// fmt.Println(*s2)

	// closures()

	// multi_goroutine()

	// goroutine_channel()
	// goroutine_channel_2()

	// for {
	// 	_, err := Get("http://www.baidu.com/")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// }

	// go utils.Foo()
	// fmt.Println("打印4")
	// time.Sleep(1000 * time.Second)
}
